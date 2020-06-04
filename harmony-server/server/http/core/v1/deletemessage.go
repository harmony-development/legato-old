package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"harmony-server/util"

	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// DeleteMessage deletes a message
func (h Handlers) DeleteMessage(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)

	isOwner, err := h.Deps.DB.IsOwner(*ctx.Location.GuildID, ctx.UserID)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to check permissions")
	}
	if !(isOwner || ctx.Location.Message.UserID == ctx.UserID) {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete message")
	}
	if err := h.Deps.DB.DeleteMessage(*ctx.Location.GuildID, *ctx.Location.ChannelID, ctx.Location.Message.MessageID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete message, please try again later")
	}
	h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
		Type: MessageDeleteEventType,
		Data: MessageDeleteEvent{
			GuildID:   util.U64TS(*ctx.Location.GuildID),
			ChannelID: util.U64TS(*ctx.Location.ChannelID),
			MessageID: util.U64TS(ctx.Location.Message.MessageID),
		},
	})
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted message",
	})
}
