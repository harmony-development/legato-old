package v1

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"harmony-server/server/db/queries"
	"harmony-server/server/http/responses"
	"harmony-server/util"

	"harmony-server/server/http/hm"

	"github.com/labstack/echo/v4"
)

// GetMessagesData is the data for a message listing request
type GetMessagesData struct {
	// MessageRef is the ID of the message you want to load before.
	// Used to load old messages
	MessageRef string `json:"before_message"`
}

// GetMessages gets messages in a given channel
func (h Handlers) GetMessages(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(GetMessagesData)
	var messageRef uint64
	if data.MessageRef != "" {
		var err error
		messageRef, err = strconv.ParseUint(data.MessageRef, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
	}

	exists, err := h.Deps.DB.UserInGuild(ctx.UserID, *ctx.Location.GuildID)
	if err != nil || !exists {
		return echo.NewHTTPError(http.StatusForbidden, "not allowed to get messages")
	}
	var messages []queries.Message
	if messageRef == 0 {
		messages, err = h.Deps.DB.GetMessages(*ctx.Location.GuildID, *ctx.Location.ChannelID)
	} else {
		time, err := h.Deps.DB.GetMessageDate(messageRef)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "error getting message date")
		}
		messages, err = h.Deps.DB.GetMessagesBefore(*ctx.Location.GuildID, *ctx.Location.ChannelID, time)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, responses.UnknownError)
		}
	}
	if err != nil && err != sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusInternalServerError, "error listing messages")
	}
	return ctx.JSON(http.StatusOK, MessageListResponse{
		Messages: func() []Message {
			var ret []Message
			for _, message := range messages {
				var embeds []Embed
				var actions []Action
				for _, rawEmbed := range message.Embeds {
					var embed Embed
					if err := json.Unmarshal(rawEmbed, &embed); err != nil {
						continue
					}
					embeds = append(embeds, embed)
				}
				for _, rawAction := range message.Actions {
					var action Action
					if err := json.Unmarshal(rawAction, &action); err != nil {
						continue
					}
					actions = append(actions, action)
				}
				ret = append(ret, Message{
					MessageID: util.U64TS(message.MessageID),
					GuildID:   util.U64TS(message.GuildID),
					ChannelID: util.U64TS(message.ChannelID),
					AuthorID:  util.U64TS(message.UserID),
					CreatedAt: util.TimeTS(message.CreatedAt),
					EditedAt:  util.NullTimeTS(message.EditedAt),
					Content:   message.Content,
					Embeds:    embeds,
					Actions:   actions,
				})
			}
			return ret
		}(),
	})
}
