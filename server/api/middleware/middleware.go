package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"golang.org/x/time/rate"
)

type RateLimit struct {
	Duration time.Duration
	Burst    int
}

type RPCConfig struct {
	RateLimit RateLimit
	Auth      bool
}

var rpcConfigs = map[string]RPCConfig{
	"/protocol.profile.v1.ProfileService/GetUser": {
		RateLimit: RateLimit{
			Duration: 10 * time.Second,
			Burst:    64,
		},
		Auth: true,
	},
	"/protocol.profile.v1.ProfileService/GetUserMetadata": {
		RateLimit: RateLimit{Duration: 1 * time.Second,
			Burst: 4,
		},
		Auth: true,
	},
	"/protocol.profile.v1.ProfileService/UsernameUpdate": {
		RateLimit: RateLimit{
			Duration: 5 * time.Minute,
			Burst:    8,
		},
		Auth: true,
	},
	"/protocol.profile.v1.ProfileService/StatusUpdate": {
		RateLimit: RateLimit{
			Duration: 5 * time.Second,
			Burst:    4,
		},
		Auth: true,
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
	DB     db.IHarmonyDB
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
