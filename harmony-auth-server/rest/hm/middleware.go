package hm

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

type HarmonyContext struct {
	echo.Context
	Limiter *rate.Limiter
	UserID *string
}

type HarmonyHandler func(ctx HarmonyContext) error