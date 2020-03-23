package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/db"
	"harmony-server/globals"
	"harmony-server/rest/hm"
	"net/http"
)

func GetMembers(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	guild := ctx.FormValue("guild")
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[ctx.User.ID] == nil {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to list members")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many member listing requests, please try again later")
	}
	res, err := db.DBInst.Query("SELECT userid FROM guildmembers WHERE guildid=$1", guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to list members, please try again later")
	}
	var returnMembers []string
	for res.Next() {
		var userid string
		err = res.Scan(&userid)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to list members, please try again later")
		}
		returnMembers = append(returnMembers, userid)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"members": returnMembers,
	})
}
