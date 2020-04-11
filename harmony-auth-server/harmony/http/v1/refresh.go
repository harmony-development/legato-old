package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/harmony/http/hm"
	"net/http"
	"time"
)

type refreshData struct {
	Session string `validate:"required"`
}

// Refresh extends the lifespan of a token
func (h Handlers) Refresh(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := new(refreshData)
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many requests, please try again later")
	}
	session, exists := h.AuthManager.Sessions.GetSession(data.Session)
	if !exists {
		return echo.NewHTTPError(http.StatusUnauthorized, "session required to refresh")
	}
	session.Expiration = time.Now().Add(h.Config.Server.SessionExpire)
	return ctx.NoContent(http.StatusOK)
}
