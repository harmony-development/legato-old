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
	ActionID string `validate:"required"`
	Data     string
}

// SendActionData is the data that will be sent to a client
type SendActionData struct {
	GuildID   uint64 `json:"guildID"`
	ChannelID uint64 `json:"channelID"`
	MessageID uint64 `json:"messageID"`
	TriggerID uint64 `json:"triggerID"`
	Action    struct {
		ID   string `json:"id"`
		Data string `json:"data"`
	} `json:"action"`
}

// TriggerAction will trigger an action for a client to receive
func (h Handlers) TriggerAction(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(TriggerActionData)

	handle := func() *guild.ClientArray {
		for id, client := range h.Deps.State.Guilds[*ctx.Location.GuildID].Clients {
			if id == ctx.Location.Message.UserID {
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
				GuildID:   *ctx.Location.GuildID,
				ChannelID: ctx.Location.Message.ChannelID,
				MessageID: ctx.Location.Message.MessageID,
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
