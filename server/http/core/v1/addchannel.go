package v1

import (
	"net/http"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/socket/client"
	"github.com/harmony-development/legato/util"

	"github.com/labstack/echo/v4"
)

// AddChannelData represents data received from client on AddChannel
type AddChannelData struct {
	ChannelName string `json:"channel_name" validate:"required"`
}

// AddChannel is a request to add a channel to a guild
func (h Handlers) AddChannel(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(AddChannelData)

	channel, err := h.Deps.DB.AddChannelToGuild(*ctx.Location.GuildID, data.ChannelName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	if h.Deps.State.Guilds[*ctx.Location.GuildID] != nil {
		h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
			Type: ChannelCreateEventType,
			Data: ChannelCreateEvent{
				GuildID:     util.U64TS(*ctx.Location.GuildID),
				ChannelName: data.ChannelName,
				ChannelID:   util.U64TS(channel.ChannelID),
			},
		})
	}
	return ctx.JSON(http.StatusOK, ChannelCreateResponse{
		GuildID:     util.U64TS(*ctx.Location.GuildID),
		ChannelName: data.ChannelName,
		ChannelID:   util.U64TS(channel.ChannelID),
	})
}
