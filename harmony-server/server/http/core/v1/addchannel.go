package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AddChannelData represents data received from client on AddChannel
type AddChannelData struct {
	Guild       int64  `validate:"required"`
	ChannelName string `validate:"required"`
}

// AddChannel is a request to add a channel to a guild
func (h Handlers) AddChannel(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data AddChannelData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	if err := ctx.VerifyOwner(h.Deps.DB, data.Guild, ctx.UserID); err != nil {
		return err
	}
	channel, err := h.Deps.DB.AddChannelToGuild(data.Guild, data.ChannelName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
		Type: "AddChannel",
		Data: map[string]interface{}{
			"guild":       data.Guild,
			"channelName": data.ChannelName,
			"channelID":   channel.ChannelID,
		},
	})
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully added channel",
	})
}
