package v1

import (
	"encoding/json"
	"net/http"

	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"harmony-server/util"

	"github.com/labstack/echo/v4"
)

type MessageUpdateData struct {
	Content *string
	Embeds  *[]string
	Actions *[]string
}

func (h Handlers) UpdateMessage(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(MessageUpdateData)

	if ctx.UserID != ctx.Location.Message.UserID {
		return echo.NewHTTPError(http.StatusForbidden)
	}
	if data.Content == nil && data.Embeds == nil && data.Actions == nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	var actions, embeds [][]byte
	var ptrActions, ptrEmbeds *[][]byte
	var rawEmbeds, rawActions []json.RawMessage
	var flags MessageUpdateFlags
	if data.Content != nil {
		flags |= UpdateContent
	}
	if data.Actions != nil && len(*data.Actions) > 0 {
		for _, action := range *data.Actions {
			parsed, err := CleanAction([]byte(action))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			actions = append(actions, parsed)
		}
		for _, action := range actions {
			rawActions = append(rawActions, json.RawMessage(action))
		}
		flags |= UpdateActions
		ptrActions = &actions
	}
	if data.Embeds != nil && len(*data.Embeds) > 0 {
		for _, embed := range *data.Embeds {
			parsed, err := CleanEmbed([]byte(embed))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			embeds = append(embeds, parsed)
		}
		for _, embed := range embeds {
			rawEmbeds = append(rawEmbeds, json.RawMessage(embed))
		}
		flags |= UpdateEmbeds
		ptrEmbeds = &embeds
	}
	editedAt, err := h.Deps.DB.UpdateMessage(ctx.Location.Message.MessageID, data.Content, ptrEmbeds, ptrActions)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
		Type: MessageUpdateEventType,
		Data: MessageUpdateEvent{
			GuildID:   util.u64TS(*ctx.Location.GuildID),
			ChannelID: util.u64TS(*ctx.Location.ChannelID),
			MessageID: util.u64TS(ctx.Location.Message.MessageID),
			Flags:     flags,
			EditedAt:  editedAt.Unix(),
			Message:   *data.Content,
			Actions:   rawActions,
			Embeds:    rawEmbeds,
		},
	})
	return nil
}
