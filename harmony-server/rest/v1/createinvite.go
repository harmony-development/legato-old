package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func CreateInvite(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	guild := ctx.FormValue("guildid")
	token := ctx.FormValue("token")
	userid, err := authentication.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many invites created, try again in a few seconds")
	}
	var inviteID = randstr.Hex(5)
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Owner != userid {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to create an invite")
	}
	_, err = harmonydb.DBInst.Exec("INSERT INTO invites(inviteid, guildid) VALUES($1, $2)", inviteID, guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating invite, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"invite": inviteID,
	})
}