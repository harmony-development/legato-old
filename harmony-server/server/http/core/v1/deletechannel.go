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
	Guild   int64 `validate:"required"`
	Channel int64 `validate:"required"`
}

// DeleteChannel is the request to delete a channel
func (h Handlers) DeleteChannel(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data DeleteChannelData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	if err := ctx.VerifyOwner(h.Deps.DB, data.Guild, ctx.UserID); err != nil {
		return err
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
