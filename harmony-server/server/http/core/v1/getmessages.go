package v1

import (
	"harmony-server/server/db/queries"
	"net/http"

	"harmony-server/server/http/hm"

	"github.com/labstack/echo/v4"
)

// GetMessagesData is the data for a message listing request
type GetMessagesData struct {
	// MessageRef is the ID of the message you want to load before.
	// Used to load old messages
	MessageRef uint64
}

// GetMessages gets messages in a given channel
func (h Handlers) GetMessages(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(GetMessagesData)

	exists, err := h.Deps.DB.UserInGuild(ctx.UserID, *ctx.Location.GuildID)
	if err != nil || !exists {
		return echo.NewHTTPError(http.StatusForbidden, "not allowed to get messages")
	}
	var messages []queries.GetMessagesRow
	if data.MessageRef != 0 {
		messages, err = h.Deps.DB.GetMessages(*ctx.Location.GuildID, *ctx.Location.ChannelID)
	} else {
		time, err := h.Deps.DB.GetMessageDate(data.MessageRef)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error getting message date")
		}
		messages, err = h.Deps.DB.GetMessagesBefore(*ctx.Location.GuildID, *ctx.Location.ChannelID, time)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error listing messages")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": messages,
	})
}
