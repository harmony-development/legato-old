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

// TriggerAction will trigger an action for a client to receive
func (h Handlers) TriggerAction(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(TriggerActionData)

	handle := func() *guild.ClientArray {
		if val, ok := h.Deps.State.Guilds[*ctx.Location.GuildID].Clients[ctx.Location.Message.MessageID]; ok {
			return val
		}
		return nil
	}()
	if handle == nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	for _, conn := range handle.Clients {
		conn.Send(&client.OutPacket{
			Type: ActionEventType,
			Data: ActionEvent{
				GuildID:   u64TS(*ctx.Location.GuildID),
				ChannelID: u64TS(ctx.Location.Message.ChannelID),
				MessageID: u64TS(ctx.Location.Message.MessageID),
				TriggerID: u64TS(ctx.UserID),
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
