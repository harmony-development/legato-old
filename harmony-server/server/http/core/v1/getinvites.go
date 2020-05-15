package v1

import (
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"harmony-server/server/http/hm"
	"net/http"
)

type GetInvitesData struct {
	Guild string `validate:"required"`
}

func (h Handlers) GetInvites(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(*GetInvitesData)

	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] == nil {
		return echo.NewHTTPError(http.StatusNotFound, "guild not found")
	}
	if isOwner, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID); err != nil || !isOwner {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permission to list invites")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many invite listing requests, please try again later")
	}
	invites, err := h.Deps.DB.GetInvites(data.Guild)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to get invites")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"invites": invites,
	})
}
