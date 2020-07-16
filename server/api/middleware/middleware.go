package middleware

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Limit struct {
	Duration time.Duration
	Burst    int
}

var limits = map[string]Limit{
	"/protocol.profile.v1.ProfileService/GetUser": {
		Duration: 10 * time.Second,
		Burst:    64,
	},
	"/protocol.profile.v1.ProfileService/GetUserMetadata": {
		Duration: 1 * time.Second,
		Burst:    4,
	},
	"/protocol.profile.v1.ProfileService/UsernameUpdate": {
		Duration: 5 * time.Minute,
		Burst:    8,
	},
	"/protocol.profile.v1.ProfileService/StatusUpdate": {
		Duration: 5 * time.Second,
		Burst:    4,
	},
}

// HarmonyContext contains a custom context for passing data from middleware to handlers
type HarmonyContext struct {
	context.Context
	UserID  uint64
	Limiter *rate.Limiter
}

type Middlewares struct {
	RateLock *sync.RWMutex
	// RateLimits is a map of IP rate limits for each RPC route
	RateLimits map[string]map[string]visitor
}
