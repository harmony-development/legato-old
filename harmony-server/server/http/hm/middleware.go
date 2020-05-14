package hm

import (
	"sync"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"harmony-server/server/db"
)

type HarmonyContext struct {
	echo.Context
	Limiter *rate.Limiter
	UserID  string
}

type HarmonyHandler func(ctx HarmonyContext) error

type Middlewares struct {
	DB *db.DB
	RateLimits map[string]map[string]*visitor
	RateLock sync.RWMutex
}

func New(db *db.DB) *Middlewares {
	m := &Middlewares{
		DB: db,
	}
	go m.RateCleanup()
	return m
}