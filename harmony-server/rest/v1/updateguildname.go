package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/globals"
	"harmony-server/db"
	"harmony-server/rest/hm"
	"net/http"
)

type updateGuildName struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
	Name string `mapstructure:"name"`
}

func UpdateGuildName(c echo.Context) error {
	ctx := c.(*hm.HarmonyContext)
	guild, name := ctx.FormValue("guild"), ctx.FormValue("name")
	if globals.Guilds[guild] == nil || globals.Guilds[guild].Clients[ctx.User.ID] == nil || globals.Guilds[guild].Owner != ctx.User.ID {
		return echo.NewHTTPError(http.StatusForbidden, "insufficient perms to update guild name")
	}
	if !ctx.Limiter.Allow() {
		return echo.NewHTTPError(http.StatusTooManyRequests, "too many guild name update requests, please try again later")
	}
	_, err := db.DBInst.Exec("UPDATE guilds SET guildname=$1 WHERE guildid=$2", name, guild)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update guild name, please try again later")
	}
	for _, client := range globals.Guilds[guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "updateguildname",
				Data: map[string]interface{}{
					"guild": guild,
					"name":  name,
				},
			})
		}
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "successfully updated guild name",
	})
}