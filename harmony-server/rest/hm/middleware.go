package hm

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
)

type HarmonyContext struct {
	echo.Context
	Limiter *rate.Limiter
	User *authentication.SessionData
}

type HarmonyHandler func(ctx HarmonyContext) error