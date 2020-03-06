package v1

import (
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket/event"
)

type updateGuildName struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
	Name string `mapstructure:"name"`
}

func UpdateGuildName(limiter *rate.Limiter, c echo.Context) error {
	var data updateGuildName
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil || globals.Guilds[data.Guild].Owner != ws.Userid {
		return
	}
	if !ctx.Limiter.Allow() {
		event.sendErr(ws, "You're updating the guild name a bit too fast... try again in a few seconds")
		return
	}
	_, err := harmonydb.DBInst.Exec("UPDATE guilds SET guildname=$1 WHERE guildid=$2", data.Name, data.Guild)
	if err != nil {
		golog.Warnf("Error updating name. %v", err)
		ws.Send(&globals.Packet{
			Type: "updateguildname",
			Data: map[string]interface{}{
				"success": false,
			},
		})
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "updateguildname",
				Data: map[string]interface{}{
					"guild": data.Guild,
					"name":  data.Name,
				},
			})
		}
	}
}