package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type deleteChannelData struct {
	Token       string `mapstructure:"token"`
	Guild       string `mapstructure:"guild"`
	ChannelID   string `mapstructure:"channel"`
}

func OnDeleteChannel(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data deleteChannelData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil || globals.Guilds[data.Guild].Owner != ws.Userid {
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're deleting a lot of channels! Try again later!")
		return
	}
	transaction, err := harmonydb.DBInst.Begin()
	if err != nil {
		sendErr(ws, "We weren't able to delete that channel for some reason. You should try again")
		golog.Warnf("Error making channel delete transaction : %v", err)
		return
	}
	_, err = transaction.Exec("DELETE FROM messages WHERE channelid=$1 AND guildid=$2", data.ChannelID, data.Guild)
	if err != nil {
		sendErr(ws, "For some reason we couldn't delete the messages in that channel. You should try again")
		golog.Warnf("Error deleting channel : %v", err)
		return
	}
	_, err = transaction.Exec("DELETE FROM channels WHERE channelid=$1 AND guildid=$2", data.ChannelID, data.Guild)
	if err != nil {
		sendErr(ws, "We weren't able to delete that channel for some reason. You should try again")
		golog.Warnf("Error deleting channel : %v", err)
		return
	}
	if err = transaction.Commit(); err != nil {
		sendErr(ws, "We weren't able to delete that channel for some reason. You should try again")
		golog.Warnf("Error deleting channel : %v", err)
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		for _, conn := range client {
			conn.Send(&globals.Packet{
				Type: "deletechannel",
				Data: map[string]interface{}{
					"guild":     data.Guild,
					"channelid": data.ChannelID,
				},
			})
		}
	}
}
