package v1

import (
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"harmony-server/server/http/hm"
	"net/http"
)

type LeaveGuildData struct {
	Guild string `validate:"required"`
}

func (h Handlers) LeaveGuild(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data LeaveGuildData
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] == nil {
		return echo.NewHTTPError(http.StatusNotFound, "guild not found")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many guild leave requests, please try again later")
	}
	if isOwner, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID); err != nil || !isOwner {
		return echo.NewHTTPError(http.StatusForbidden, "you cannot leave a guild you own")
	}
	if err := h.Deps.DB.DeleteMember(data.Guild, ctx.UserID); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to leave guild, this issue has been reported")
	}
	h.Deps.State.GuildsLock.Lock()
	delete(h.Deps.State.Guilds[data.Guild].Clients, ctx.UserID)
	h.Deps.State.GuildsLock.Unlock()
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully left guild",
	})
}
