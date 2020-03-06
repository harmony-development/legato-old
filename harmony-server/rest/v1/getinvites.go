package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func GetInvites(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	guild := ctx.FormValue("guild")
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[*ctx.UserID] == nil || globals.Guilds[guild].Owner != *ctx.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to list invites")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many invite listing requests, please try again later")
	}
	res, err := harmonydb.DBInst.Query("SElECT inviteid, invitecount FROM invites WHERE guildid=$1 ORDER BY invitecount", guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list invites, please try again later")
	}
	returnInvites := make(map[string]int)
	for res.Next() {
		var invitecode string
		var invitecount int
		err = res.Scan(&invitecode, &invitecount)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "unable to get invite, please try again later")
		}
		returnInvites[invitecode] = invitecount
	}
	return ctx.JSON(http.StatusOK, map[string]map[string]int{
		"invites": returnInvites,
	})
}
