package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/db"
	"harmony-server/globals"
	"harmony-server/rest/hm"
	"net/http"
)

func DeleteInvite(c echo.Context) error {
	ctx, _ := c.(hm.HarmonyContext)
	guild, invite := ctx.FormValue("guild"), ctx.FormValue("invite")
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[ctx.User.ID] == nil || globals.Guilds[guild].Owner != ctx.User.ID {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete invite")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many invite deletions, please wait a few moments")
	}
	_, err := db.DBInst.Exec("DELETE FROM invites WHERE inviteid=$1 AND guildid=$2", invite, guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to delete invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted invite",
	})
}
