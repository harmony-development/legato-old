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

type deleteMessageData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
	Channel string `mapstructure:"channel"`
	Message string `mapstructure:"message"`
}

func DeleteMessage(limiter *rate.Limiter, ctx echo.Context) error {
	var data deleteMessageData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		golog.Warnf("Error decoding data for getting channels")
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil {
		return
	}
	if !limiter.Allow() {
		event.sendErr(ws, "You're deleting messages too fast, try again in a few moments")
		return
	}
	res, err := harmonydb.DBInst.Exec("DELETE FROM messages WHERE guildid=$1 AND channelid=$2 AND messageid=$3 AND (author=$4 OR (SELECT owner FROM guilds WHERE guildid=$5)=$6)", data.Guild, data.Channel, data.Message, ws.Userid, data.Guild, ws.Userid)
	if err != nil {
		event.sendErr(ws, "An error occured while deleting that message")
		return
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		event.sendErr(ws, "An error occured while deleting that message")
		return
	}
	if rowCount != 1 {
		event.sendErr(ws, "An error occured while deleting that message")
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "deletemessage",
				Data: map[string]interface{}{
					"guild":     data.Guild,
					"channel":   data.Channel,
					"message": data.Message,
				},
			})
		}
	}
}