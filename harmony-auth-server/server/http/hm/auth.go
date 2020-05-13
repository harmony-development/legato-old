package hm

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// WithAuth adds authentication to an endpoint
func (m *Middlewares) WithAuth(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(HarmonyContext)
		authorization := ctx.Request().Header.Get("Authorization")
		if authorization == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid-session")
		}
		session, exists := m.AuthManager.Sessions.GetSession(authorization)
		if !exists {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid-session")
		}
		ctx.Session = session
		return handler(ctx)
	}
}