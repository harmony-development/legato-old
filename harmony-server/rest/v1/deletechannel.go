package v1

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func DeleteChannel(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	token, guild, channelid := ctx.FormValue("token"), ctx.FormValue("guild"), ctx.FormValue("channelid")
	userid, err := authentication.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Owner != userid {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete channel")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many channel deletions, please try again in a few seconds")
	}
	err = harmonydb.DeleteChannelTransaction(guild, channelid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting channel, please try again later")
	}
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "deletechannel",
				Data: map[string]interface{}{
					"guild":     guild,
					"channelid": channelid,
				},
			})
		}
	}
}
