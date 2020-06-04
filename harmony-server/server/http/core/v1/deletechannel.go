package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"harmony-server/util"

	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// DeleteChannel is the request to delete a channel
func (h Handlers) DeleteChannel(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)

	if err := h.Deps.DB.DeleteChannelFromGuild(*ctx.Location.GuildID, *ctx.Location.ChannelID); err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting channel, please try again later")
	}
	if h.Deps.State.Guilds[*ctx.Location.GuildID] != nil {
		h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
			Type: ChannelDeleteEventType,
			Data: ChannelDeleteEvent{
				GuildID:   util.u64TS(*ctx.Location.GuildID),
				ChannelID: util.u64TS(*ctx.Location.ChannelID),
			},
		})
	}
	return nil
}
