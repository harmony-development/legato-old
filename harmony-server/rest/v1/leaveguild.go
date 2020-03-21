package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/globals"
	"harmony-server/db"
	"harmony-server/rest/hm"
	"net/http"
)

func LeaveGuild(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	guild := ctx.FormValue("guild")
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[ctx.User.ID] == nil {
		return echo.NewHTTPError(http.StatusForbidden, "you are not part of the guild, so you cannot leave it")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many guild leave requests, please try again later")
	}
	var ownerID  string
	err := db.DBInst.QueryRow("SELECT owner FROM guilds WHERE guildid=$1", guild).Scan(&ownerID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to leave guild, please try again later")
	}
	if ownerID == ctx.User.ID {
		return echo.NewHTTPError(http.StatusBadRequest, "you can't leave a guild you own")
	}
	_, err = db.DBInst.Exec("DELETE FROM guildmembers WHERE userid=$1 AND guildid=$2", ctx.User.ID, guild)
	// GUILD STUCK! GUILD STUCK! PLEASE! I BEG YOU!
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to leave guild, please try again later. If this persists, please submit an issue on Github")
	}
	delete(globals.Guilds[guild].Clients, ctx.User.ID)
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully left guild",
	})
}
