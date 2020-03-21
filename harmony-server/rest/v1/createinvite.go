package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"harmony-server/db"
	"harmony-server/globals"
	"harmony-server/rest/hm"
	"net/http"
)

func CreateInvite(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	guild := ctx.FormValue("guildid")
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many invites created, try again in a few seconds")
	}
	var inviteID = randstr.Hex(5)
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Owner != ctx.User.ID {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to create an invite")
	}
	_, err := db.DBInst.Exec("INSERT INTO invites(inviteid, guildid) VALUES($1, $2)", inviteID, guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"invite": inviteID,
	})
}