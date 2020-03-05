package v1

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"net/http"
)

func DeleteMessage(limiter *rate.Limiter, ctx echo.Context) error {
	token, guild, channel, message := ctx.FormValue("token"), ctx.FormValue("guild"), ctx.FormValue("channel"), ctx.FormValue("message")
	userid, err := authentication.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "")
	}
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[guild] == nil {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete message")
	}
	if !limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many message deletions, please try again later")
	}
	err = harmonydb.DeleteMessageTransaction(guild, channel, message, userid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete message, please try again later")
	}
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "deletemessage",
				Data: map[string]interface{}{
					"guild":     guild,
					"channel":   channel,
					"message": message,
				},
			})
		}
	}
}