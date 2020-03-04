package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"net/http"
)

func AddChannel(limiter *rate.Limiter, ctx echo.Context) error {
	guild := ctx.QueryParam("guildid")
	token, channelname := ctx.FormValue("token"), ctx.FormValue("channelname")
	userid, err := authentication.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	if channelname == "" || guild == "" || userid == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameters")
	}
	if !limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many channels being added, please wait a few seconds")
	}
	var channelID = randstr.Hex(16)
	_, err = harmonydb.DBInst.Exec("INSERT INTO channels(channelid, guildid, channelname) VALUES($1, $2, $3) WHERE guilds.owner=$4", channelID, guild, channelname, userid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error occurred adding the guild. Do you have sufficient permission?")
	}
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients == nil {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "successfully added channel",
		})
	} else {
		for _, client := range globals.Guilds[guild].Clients {
			for _, conn := range client {
				conn.Send(&globals.Packet{
					Type: "addchannel",
					Data: map[string]interface{}{
						"guild":       guild,
						"channelname": guild,
						"channelid":   channelID,
					},
				})
			}
			return ctx.JSON(http.StatusOK, map[string]string{
				"message": "successfully added channel",
			})
		}
	}
}