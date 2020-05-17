package v1

import (
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket/client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DeleteGuildData is the data for a guild deletion request
type DeleteGuildData struct {
	Guild uint64 `validate:"required"`
}

// DeleteGuild is the handler for a delete guild request
func (h Handlers) DeleteGuild(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	var data DeleteGuildData
	if err := ctx.BindAndVerify(&data); err != nil {
		return err
	}
	if err := ctx.VerifyOwner(h.Deps.DB, data.Guild, ctx.UserID); err != nil {
		return err
	}
	if h.Deps.DB.DeleteGuild(data.Guild) != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting guild, please try again later or report to the administrator")
	}

	h.Deps.State.GuildsLock.Lock()
	h.Deps.State.Guilds[data.Guild].Broadcast(&client.OutPacket{
		Type: "deleteGuild",
		Data: map[string]interface{}{
			"guild": data.Guild,
		},
	})
	delete(h.Deps.State.Guilds, data.Guild)
	h.Deps.State.GuildsLock.RUnlock()

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted guild",
	})
}
