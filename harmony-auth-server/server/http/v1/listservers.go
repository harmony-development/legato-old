package v1

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"

	"harmony-auth-server/server/http/hm"
)

// ListInstances returns an array of servers saved in the DB
func (h Handlers) ListInstances(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)

	if !ctx.Limiter.Reserve().OK() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many instance listing requests")
	}

	servers, err := h.DB.GetInstanceList(ctx.Session.UserID)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unknown-error")
	}
	ctx.Limiter.Allow()
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"servers": servers,
	})
}
