package v1

import (
	"net/http"

	"harmony-server/server/db"
	"harmony-server/server/http/hm"

	"github.com/labstack/echo/v4"
)

// GetMessagesData is the data for a message listing request
type GetMessagesData struct {
	Guild   string `validate:"required"`
	Channel string `validate:"required"`
	// MessageRef is the ID of the message you want to load before.
	// Used to load old messages
	MessageRef string
}

// GetMessages gets messages in a given channel
func (h Handlers) GetMessages(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(*GetMessagesData)
	exists, err := h.Deps.DB.UserInGuild(ctx.UserID, data.Guild)
	if err != nil || !exists {
		return echo.NewHTTPError(http.StatusForbidden, "not allowed to get messages")
	}
	var messages []db.Message
	if data.MessageRef != "" {
		messages, err = h.Deps.DB.GetMessages(data.Guild, data.Channel)
	} else {
		time, err := h.Deps.DB.GetMessageDate(data.MessageRef)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error getting message date")
		}
		messages, err = h.Deps.DB.GetMessagesBefore(data.Guild, data.Channel, *time)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error listing messages")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": messages,
	})
}
