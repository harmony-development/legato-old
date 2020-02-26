package rest

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

// code adapted from https://www.alexedwards.net/blog/how-to-rate-limit-http-requests

type visitor struct {
	limiter rate.Limiter
	lastSeen time.Time
}

var rateLimits = make(map[string]map[string]*visitor)
var rateLock sync.RWMutex

func CleanupRoutine() {
	for {
		time.Sleep(3 * time.Minute)
		rateLock.Lock()
		defer rateLock.Unlock()
		for _, path := range rateLimits {
			for ip, v := range path {
				if time.Now().Sub(v.lastSeen) > 3*time.Minute {
					delete(path, ip)
				}
			}
		}
	}
}

func AddRateLimit(path string) {
	rateLimits[path] = make(map[string]*visitor)
}

func getVisitor(path string, ip string) *rate.Limiter {
	rateLock.RLock()
	defer rateLock.RUnlock()
	if _, exists := rateLimits[path]; !exists {
		limiter := rate.NewLimiter(3, 3)
		rateLimits[path] = make(map[string]*visitor)
		rateLimits[path][ip] = &visitor{
			limiter:  *limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}
	if v, exists := rateLimits[path][ip]; exists {
		return &v.limiter
	} else {
		limiter := rate.NewLimiter(3, 3)
		rateLimits[path][ip] = &visitor{
			limiter:  *limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}
}