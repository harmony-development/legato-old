package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type deleteChannelData struct {
	Token       string `mapstructure:"token"`
	Guild       string `mapstructure:"guild"`
	ChannelID   string `mapstructure:"channel"`
}

func OnDeleteChannel(ws *globals.Client, rawMap map[string]interface{}) {
	var data deleteChannelData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid, err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil || globals.Guilds[data.Guild].Owner != userid {
		return
	}
	transaction, err := harmonydb.DBInst.Begin()
	if err != nil {
		golog.Warnf("Error making channel delete transaction : %v", err)
		return
	}
	_, err = transaction.Exec("DELETE FROM messages WHERE channelid=$1 AND guildid=$2", data.ChannelID, data.Guild)
	if err != nil {
		golog.Warnf("Error deleting channel : %v", err)
		return
	}
	_, err = transaction.Exec("DELETE FROM channels WHERE channelid=$1 AND guildid=$2", data.ChannelID, data.Guild)
	if err != nil {
		golog.Warnf("Error deleting channel : %v", err)
		return
	}
	if err = transaction.Commit(); err != nil {
		golog.Warnf("Error deleting channel : %v", err)
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		client.Send(&globals.Packet{
			Type: "deleteguildchannel",
			Data: map[string]interface{}{
				"guild":   data.Guild,
				"channelid": data.ChannelID,
				"success": true,
			},
		})
	}
}
