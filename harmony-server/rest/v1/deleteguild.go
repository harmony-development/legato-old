package v1

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"net/http"
)

func DeleteGuild(limiter *rate.Limiter, ctx echo.Context) error {
	token, guild := ctx.FormValue("token"), ctx.FormValue("guild")
	userid, err := authentication.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
	}
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Owner != userid {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete guild")
	}
	if !limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many guild deletions, please wait a few moments")
	}
	if harmonydb.DeleteGuildTransaction(guild) != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting guild, please try again later or report to the administrator")
	}
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "deleteguild",
				Data: map[string]interface{}{
					"guild": guild,
				},
			})
		}
	}
	delete(globals.Guilds, guild)
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted guild",
	})
}