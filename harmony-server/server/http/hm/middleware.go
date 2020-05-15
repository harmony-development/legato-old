package hm

import (
	"sync"

	"harmony-server/server/db"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

// A HarmonyContext adds rate limiting and a user ID to an echo.Context
type HarmonyContext struct {
	echo.Context
	Limiter *rate.Limiter
	UserID  string
}

// Middlewares contains middlewares for Harmony
type Middlewares struct {
	DB         *db.DB
	RateLimits map[string]map[string]*visitor
	RateLock   sync.RWMutex
}

// New instantiates the middlewares for Harmony
func New(db *db.DB) *Middlewares {
	m := &Middlewares{
		DB:         db,
		RateLimits: make(map[string]map[string]*visitor),
	}
	go m.RateCleanup()
	return m
}
