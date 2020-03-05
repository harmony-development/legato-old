package v1

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/harmonydb"
	"net/http"
)

func CreateGuild(limiter *rate.Limiter, ctx echo.Context) error {
	token, guildname := ctx.FormValue("token"), ctx.FormValue("guildname")
	userid, err := authentication.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	if !limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "you're creating too many guilds, please try again in a minute or two")
	}
	guildid, err := harmonydb.CreateGuildTransaction(guildname, userid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "unable to create guild, please try again later")
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"guild": *guildid,
	})
}