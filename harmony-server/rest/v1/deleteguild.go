package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/globals"
	"harmony-server/db"
	"harmony-server/rest/hm"
	"net/http"
)

func DeleteGuild(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	guild := ctx.FormValue("guild")
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Owner != ctx.User.ID {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete guild")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many guild deletions, please wait a few moments")
	}
	if db.DeleteGuildTransaction(guild) != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting guild, please try again later or report to the administrator")
	}
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "deleteguild",
				Data: map[string]interface{}{
					"guild": guild,
				},
			})
		}
	}
	delete(globals.Guilds, guild)
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted guild",
	})
}