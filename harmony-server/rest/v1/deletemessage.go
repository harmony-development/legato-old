package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/rest/hm"
	"net/http"
)

func DeleteMessage(c echo.Context) error {
	ctx, _ := c.(*hm.HarmonyContext)
	guild, channel, message := ctx.FormValue("guild"), ctx.FormValue("channel"), ctx.FormValue("message")
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[guild] == nil {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions to delete message")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many message deletions, please try again later")
	}
	err := harmonydb.DeleteMessageTransaction(guild, channel, message, *ctx.UserID)
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
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted message",
	})
}