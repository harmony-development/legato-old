package v1

import (
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"
)

type DeleteGuildData struct {
	Guild string `validate:"required"`
}

func (h Handlers) DeleteGuild(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data DeleteGuildData
	if err := ctx.Bind(data); err != nil {
		return err
	}
	if err := ctx.Validate(data); err != nil {
		return err
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] == nil {
		return echo.NewHTTPError(http.StatusNotFound, "guild not found")
	}
	isOwner, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error checking permissions")
	}
	if !isOwner {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete guild")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many guild deletions, please wait a few moments")
	}
	if h.Deps.DB.DeleteGuild(data.Guild) != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting guild, please try again later or report to the administrator")
	}
	h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
		Type: "deleteGuild",
		Data: map[string]interface{}{
			"guild": data.Guild,
		},
	})
	h.Deps.State.GuildsLock.Lock()
	delete(h.Deps.State.Guilds, data.Guild)
	h.Deps.State.GuildsLock.RUnlock()
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted guild",
	})
}