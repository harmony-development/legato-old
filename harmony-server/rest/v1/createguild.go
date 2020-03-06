package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func CreateGuild(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	guildname := ctx.FormValue("guildname")
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "you're creating too many guilds, please try again in a minute or two")
	}
	guildid, err := harmonydb.CreateGuildTransaction(guildname, *ctx.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to create guild, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"guild": *guildid,
	})
}