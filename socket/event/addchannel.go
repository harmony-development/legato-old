package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"github.com/thanhpk/randstr"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type addChannelData struct {
	Token   string `mapstructure:"token"`
	Guild   string `mapstructure:"guild"`
	Channel string `mapstructure:"channel"`
}

func OnAddChannel(ws *globals.Client, rawMap map[string]interface{}) {
	var data addChannelData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if data.Token == "" {
		deauth(ws)
		return
	}
	if data.Guild == "" || data.Channel == "" {
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
	var channelID = randstr.Hex(16)
	_, err = harmonydb.DBInst.Exec("INSERT INTO channels(channelid, guildid, channelname) VALUES($1, $2, $3)", channelID, data.Guild, data.Channel)
	if err != nil {
		sendErr(ws, "Hmm the channel couldn't be created. You should try again.")
		golog.Warnf("Error creating channel : %v", err)
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		client.Send(&globals.Packet{
			Type: "addguildchannel",
			Data: map[string]interface{}{
				"guild":       data.Guild,
				"channelname": data.Channel,
				"channelid":   channelID,
				"success":     true,
			},
		})
	}
}
