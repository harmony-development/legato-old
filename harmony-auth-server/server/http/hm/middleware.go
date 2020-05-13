package hm

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"

	"harmony-auth-server/server/auth"
)

// HarmonyContext is a custom context that contains values from the middleware
type HarmonyContext struct {
	echo.Context
	Limiter *rate.Limiter
	Session *auth.Session
}

// Middlewares is an instance of the HTTP middleware provider
type Middlewares struct {
	AuthManager *auth.Manager
}
