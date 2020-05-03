package hm

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"harmony-server/server/db"
)

type HarmonyContext struct {
	echo.Context
	Limiter *rate.Limiter
	UserID string
}

type HarmonyHandler func(ctx HarmonyContext) error

type Middlewares struct {
	DB *db.DB
}