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

var rateLimits = make(map[string]*visitor)
var rateLock sync.RWMutex

func CleanupRoutine() {
	for {
		time.Sleep(3 * time.Minute)
		rateLock.Lock()
		defer rateLock.Unlock()
		for ip, v := range rateLimits {
			if time.Now().Sub(v.lastSeen) > 3*time.Minute {
				delete(rateLimits, ip)
			}
		}
	}
}

func getVisitor(ip string) *rate.Limiter {
	rateLock.RLock()
	defer rateLock.RUnlock()
	if v, exists := rateLimits[ip]; exists {
		return &v.limiter
	} else {
		limiter := rate.NewLimiter(3, 3)
		rateLimits[ip] = &visitor{
			limiter:  *limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}
}