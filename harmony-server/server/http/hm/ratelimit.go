package hm

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

// code adapted from https://www.alexedwards.net/blog/how-to-rate-limit-http-requests

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var rateLimits = make(map[string]map[string]*visitor)
var rateLock = sync.RWMutex{}

// CleanupRoutine cleans up old visitors from memory
func CleanupRoutine() {
	for {
		time.Sleep(3 * time.Minute)
		rateLock.Lock()
		for _, path := range rateLimits {
			for ip, v := range path {
				if time.Now().Sub(v.lastSeen) > 3*time.Minute {
					delete(path, ip)
				}
			}
		}
		rateLock.Unlock()
	}
}

func (m *Middlewares) WithRateLimit(handler echo.HandlerFunc, duration time.Duration, burst int) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(HarmonyContext)
		ctx.Limiter = getVisitor(ctx.Path(), ctx.RealIP(), duration, burst)
		return handler(ctx)
	}
}

func getVisitor(path string, ip string, duration time.Duration, burst int) *rate.Limiter {
	rateLock.Lock()
	defer rateLock.Unlock()
	if _, exists := rateLimits[path]; !exists {
		limiter := rate.NewLimiter(rate.Every(duration), burst)
		rateLimits[path] = make(map[string]*visitor)
		rateLimits[path][ip] = &visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}
	if v, exists := rateLimits[path][ip]; !exists {
		limiter := rate.NewLimiter(rate.Every(duration), burst)
		rateLimits[path][ip] = &visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	} else {
		return v.limiter
	}
}
