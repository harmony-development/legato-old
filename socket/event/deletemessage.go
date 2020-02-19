package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type deleteMessageData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
	Channel string `mapstructure:"channel"`
	Message string `mapstructure:"message"`
}

func OnDeleteMessage(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data deleteMessageData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		golog.Warnf("Error decoding data for getting channels")
		return
	}
	userid, err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil {
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're deleting messages too fast, try again in a few moments")
		return
	}
	res, err := harmonydb.DBInst.Exec("DELETE FROM messages WHERE guildid=$1 AND channelid=$2 AND messageid=$3 AND author=$4", data.Guild, data.Channel, data.Message, userid)
	if err != nil {
		sendErr(ws, "An error occured while deleting that message")
		return
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		sendErr(ws, "An error occured while deleting that message")
		return
	}
	if rowCount != 1 {
		sendErr(ws, "An error occured while deleting that message")
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