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

// HarmonyHandler is a type of handler that takes a HarmonyContext
type HarmonyHandler func(ctx HarmonyContext) error

// Middlewares contains middlewares for Harmony
type Middlewares struct {
	DB         *db.HarmonyDB
	RateLimits map[string]map[string]*visitor
	RateLock   sync.RWMutex
}

// New instantiates the middlewares for Harmony
func New(db *db.HarmonyDB) *Middlewares {
	m := &Middlewares{
		DB: db,
	}
	go m.RateCleanup()
	return m
}
