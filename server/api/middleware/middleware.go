package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/harmony-development/legato/server/logger"
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

type Dependencies struct {
	Logger logger.ILogger
}

type Middlewares struct {
	Dependencies
	RateLock *sync.RWMutex
	// RateLimits is a map of IP rate limits for each RPC route
	RateLimits map[string]map[string]visitor
}

func New(deps Dependencies) Middlewares {
	return Middlewares{
		Dependencies: deps,
		RateLock:     new(sync.RWMutex),
		RateLimits:   make(map[string]map[string]visitor),
	}
}
