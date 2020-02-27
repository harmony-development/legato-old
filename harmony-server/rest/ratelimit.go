package rest

import (
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"net/http"
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

type ratedHandler func(limiter *rate.Limiter, w http.ResponseWriter, r *http.Request)

func WithRateLimit(handler ratedHandler, duration time.Duration, burst int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := mux.CurrentRoute(r).GetName()

		handler(getVisitor(path, getIP(r), duration, burst), w, r)
	}
}

func getVisitor(path string, ip string, duration time.Duration, burst int) *rate.Limiter {
	rateLock.RLock()
	defer rateLock.RUnlock()
	if _, exists := rateLimits[path]; !exists {
		limiter := rate.NewLimiter(rate.Every(duration), burst)
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
		limiter := rate.NewLimiter(rate.Every(duration), burst)
		rateLimits[path][ip] = &visitor{
			limiter:  *limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}
}