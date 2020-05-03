package hm

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Middlewares) WithAuth(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(HarmonyContext)
		session := ctx.Request().Header.Get("Authorization")
		if session == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
		}
		userID, err := m.DB.SessionToUserID(session)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
		}
		ctx.UserID = *userID
		return handler(ctx)
	}
}