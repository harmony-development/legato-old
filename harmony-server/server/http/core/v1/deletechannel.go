package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// DeleteChannelData is the data for a channel deletion request
type DeleteChannelData struct {
	Guild   string `validate:"required"`
	Channel string `validate:"required"`
}

// DeleteChannel is the request to delete a channel
func (h Handlers) DeleteChannel(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	data := ctx.Data.(DeleteChannelData)


	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] == nil {
		return echo.NewHTTPError(http.StatusNotFound, "guild not found")
	}
	exists, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID)
	if err != nil || !exists {
		if err != nil {
			sentry.CaptureException(err)
		}
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete channel")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many channel deletions, please try again in a few seconds")
	}
	if err := h.Deps.DB.DeleteChannelFromGuild(data.Guild, data.Channel); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting channel, please try again later")
	}
	h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
		Type: "DeleteChannel",
		Data: map[string]interface{}{
			"guild":     data.Guild,
			"channelID": data.Channel,
		},
	})
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted channel",
	})
}
