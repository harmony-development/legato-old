package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type deleteChannelData struct {
	Token       string `mapstructure:"token"`
	Guild       string `mapstructure:"guild"`
	ChannelID   string `mapstructure:"channel"`
}

func OnDeleteChannel(ws *socket.Client, rawMap map[string]interface{}) {
	var data deleteChannelData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil || globals.Guilds[data.Guild].Owner != userid {
		return
	}
	_, err := harmonydb.DBInst.Exec("DELETE FROM channels WHERE channelid=? AND guildid=?", data.ChannelID, data.Guild)
	if err != nil {
		golog.Warnf("Error creating channel : %v", err)
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		client.Send(&socket.Packet{
			Type: "deleteguildchannel",
			Data: map[string]interface{}{
				"guild":   data.Guild,
				"channelid": data.ChannelID,
				"success": true,
			},
		})
	}
}
