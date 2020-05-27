package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UpdateGuildNameData is the data for a guild name update request
type UpdateGuildNameData struct {
	Name string `validate:"name"`
}

// UpdateGuildName updates the guild name
func (h Handlers) UpdateGuildName(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	data := ctx.Data.(UpdateGuildNameData)

	if err := h.Deps.DB.UpdateGuildName(*ctx.Location.GuildID, data.Name); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to update guild name, please try again later")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[*ctx.Location.GuildID] != nil {
		h.Deps.State.Guilds[*ctx.Location.GuildID].Broadcast(&client.OutPacket{
			Type: GuildUpdateEventType,
			Data: GuildUpdateEvent{
				GuildID: *ctx.Location.GuildID,
				Name:    data.Name,
			},
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated guild name",
	})
}
