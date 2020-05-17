package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"harmony-server/server/state/guild"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TriggerActionData represents data received from client on TriggerAction
type TriggerActionData struct {
	GuildID   int64  `validate:"required"`
	MessageID int64  `validate:"required"`
	BotID     int64  `validate:"required"`
	ActionID  string `validate:"required"`
	Data      string
}

// SendActionData is the data that will be sent to a client
type SendActionData struct {
	GuildID   int64 `json:"guildID"`
	ChannelID int64 `json:"channelID"`
	MessageID int64 `json:"messageID"`
	TriggerID int64 `json:"triggerID"`
	Action    struct {
		ID   string `json:"id"`
		Data string `json:"data"`
	} `json:"action"`
}

// TriggerAction will trigger an action for a client to receive
func (h Handlers) TriggerAction(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data TriggerActionData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	msg, err := h.Deps.DB.GetMessage(data.MessageID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if msg.UserID != data.BotID {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	handle := func() *guild.ClientArray {
		for id, client := range h.Deps.State.Guilds[data.GuildID].Clients {
			if id == data.BotID {
				return client
			}
		}
		return nil
	}()
	if handle == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	for _, conn := range handle.Clients {
		conn.Send(&client.OutPacket{
			Type: "action",
			Data: SendActionData{
				GuildID:   data.GuildID,
				ChannelID: msg.ChannelID,
				MessageID: msg.MessageID,
				TriggerID: ctx.UserID,
				Action: struct {
					ID   string "json:\"id\""
					Data string "json:\"data\""
				}{
					ID:   data.ActionID,
					Data: data.Data,
				},
			},
		})
	}
	return echo.NewHTTPError(http.StatusOK)
}
