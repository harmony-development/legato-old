package hm

import (
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

// code adapted from https://www.alexedwards.net/blog/how-to-rate-limit-http-requests

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// RateCleanup cleans up old visitors from memory
func (m *Middlewares) RateCleanup() {
	for {
		time.Sleep(3 * time.Minute)
		m.RateLock.Lock()
		for _, path := range m.RateLimits {
			for ip, v := range path {
				if time.Now().Sub(v.lastSeen) > 3*time.Minute {
					delete(path, ip)
				}
			}
		}
		m.RateLock.Unlock()
	}
}

// RateLimit enforces a rate limit on clients
func (m *Middlewares) RateLimit(duration time.Duration, burst int) echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.(HarmonyContext)
			ctx.Limiter = m.GetVisitor(ctx.Path(), ctx.RealIP(), duration, burst)
			return handler(ctx)
		}
	}
}

// GetVisitor gets the rate limiter for a visitor
func (m *Middlewares) GetVisitor(path string, ip string, duration time.Duration, burst int) *rate.Limiter {
	m.RateLock.Lock()
	defer m.RateLock.Unlock()
	if _, exists := m.RateLimits[path]; !exists {
		m.RateLimits[path] = make(map[string]*visitor)
	}
	if v, exists := m.RateLimits[path][ip]; !exists {
		limiter := rate.NewLimiter(rate.Every(duration), burst)
		m.RateLimits[path][ip] = &visitor{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	} else {
		return v.limiter
	}
}
