package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
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
	guildid := randstr.Hex(16)
	createGuildTransaction, err := harmonydb.DBInst.Begin()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating guild, please try again later")
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO guilds(guildid, guildname, picture, owner) VALUES($1, $2, $3, $4);`, guildid, guildname, "", userid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating guild, please try again later")
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO guildmembers(userid, guildid) VALUES($1, $2);`, userid, guildid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating guild, please try again later")
	}
	_, err = createGuildTransaction.Exec(`INSERT INTO channels(channelid, guildid, channelname) VALUES($1, $2, $3)`, randstr.Hex(16), guildid, "general")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating guild, please try again later")
	}
	err = createGuildTransaction.Commit()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating guild, please try again later")
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"guild": guildid,
	})
}