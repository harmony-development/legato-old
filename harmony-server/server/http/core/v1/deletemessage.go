package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

// DeleteMessageData is the data for a message deletion request
type DeleteMessageData struct {
	Guild     int64 `validate:"guild"`
	Channel   int64 `validate:"channel"`
	MessageID int64 `validate:"message"`
}

// DeleteMessage deletes a message
func (h Handlers) DeleteMessage(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data DeleteMessageData
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] == nil {
		return echo.NewHTTPError(http.StatusNotFound, "guild not found")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many message deletions, please try again later")
	}
	isOwner, err := h.Deps.DB.IsOwner(data.Guild, ctx.UserID)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to check permissions")
	}
	messageOwner, err := h.Deps.DB.GetMessageOwner(data.MessageID)
	if err != nil {
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to check permissions")
	}
	if !(isOwner || messageOwner == ctx.UserID) {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete message")
	}
	if err := h.Deps.DB.DeleteMessage(data.Guild, data.Channel, data.MessageID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete message, please try again later")
	}
	h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
		Type: "deleteMessage",
		Data: map[string]interface{}{
			"guild":     data.Guild,
			"channel":   data.Channel,
			"messageID": data.MessageID,
		},
	})
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted message",
	})
}
