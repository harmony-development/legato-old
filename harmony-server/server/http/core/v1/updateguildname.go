package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UpdateGuildNameData is the data for a guild name update request
type UpdateGuildNameData struct {
	Guild uint64 `validate:"required"`
	Name  string `validate:"name"`
}

// UpdateGuildName updates the guild name
func (h Handlers) UpdateGuildName(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	var data UpdateGuildNameData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	if err := ctx.VerifyOwner(h.Deps.DB, data.Guild, ctx.UserID); err != nil {
		return err
	}

	if err := h.Deps.DB.UpdateGuildName(data.Guild, data.Name); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to update guild name, please try again later")
	}
	h.Deps.State.GuildsLock.RLock()
	defer h.Deps.State.GuildsLock.RUnlock()
	if h.Deps.State.Guilds[data.Guild] != nil {
		h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
			Type: "GuildNameUpdate",
			Data: map[string]interface{}{
				"guild": data.Guild,
				"name":  data.Name,
			},
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated guild name",
	})
}
