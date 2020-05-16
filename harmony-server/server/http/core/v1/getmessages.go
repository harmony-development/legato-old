package v1

import (
	"harmony-server/server/db/queries"
	"harmony-server/server/http/hm"

	"net/http"

	"github.com/labstack/echo/v4"
)

// GetMessagesData is the data for a message listing request
type GetMessagesData struct {
	Guild   int64 `validate:"required"`
	Channel int64 `validate:"required"`
	// MessageRef is the ID of the message you want to load before.
	// Used to load old messages
	MessageRef int64
}

// GetMessages gets messages in a given channel
func (h Handlers) GetMessages(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data GetMessagesData
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if err := ctx.Validate(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	exists, err := h.Deps.DB.UserInGuild(ctx.UserID, data.Guild)
	if err != nil || !exists {
		return echo.NewHTTPError(http.StatusForbidden, "not allowed to get messages")
	}
	var messages []queries.GetMessagesRow
	if data.MessageRef != 0 {
		messages, err = h.Deps.DB.GetMessages(data.Guild, data.Channel)
	} else {
		time, err := h.Deps.DB.GetMessageDate(data.MessageRef)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error getting message date")
		}
		messages, err = h.Deps.DB.GetMessagesBefore(data.Guild, data.Channel, time)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error listing messages")
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": messages,
	})
}
