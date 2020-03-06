package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func JoinGuild(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	inviteCode := ctx.FormValue("inviteCode")
	var guildid string
	err := harmonydb.DBInst.QueryRow("SELECT guildid FROM invites WHERE inviteid=$1", inviteCode).Scan(&guildid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to join guild, invite code may not exist")
	}
	if err := harmonydb.JoinGuildTransaction(*ctx.UserID, guildid, inviteCode); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to join guild, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"guild": guildid,
	})
}
