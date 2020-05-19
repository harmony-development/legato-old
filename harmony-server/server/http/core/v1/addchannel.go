package v1

import (
	"net/http"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"

	"github.com/labstack/echo/v4"
)

// AddChannelData represents data received from client on AddChannel
type AddChannelData struct {
	ChannelName string `validate:"required"`
}

// AddChannel is a request to add a channel to a guild
func (h Handlers) AddChannel(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(*AddChannelData)

	channel, err := h.Deps.DB.AddChannelToGuild(*ctx.Location.GuildID, data.ChannelName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
		Type: "AddChannel",
		Data: map[string]interface{}{
			"guild":       *ctx.Location.GuildID,
			"channelName": data.ChannelName,
			"channelID":   channel.ChannelID,
		},
	})
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully added channel",
	})
}
