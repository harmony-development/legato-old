package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type getChannelsData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnGetChannels(ws *socket.Client, rawMap map[string]interface{}) {
	var data getChannelsData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil {
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT channelid, channelname FROM channels WHERE guildid=$1", data.Guild)
	if err != nil {
		golog.Warnf("Error selecting channels : %v", err)
		return
	}

	var returnChannels = make(map[string]string)
	for res.Next() {
		var channelid, channelname string
		err = res.Scan(&channelid, &channelname)
		if err != nil {
			golog.Warnf("Error scanning channels : %v", err)
			return
		}
		returnChannels[channelid] = channelname
	}
	ws.Send(&socket.Packet{
		Type: "getchannels",
		Data: map[string]interface{}{
			"channels": returnChannels,
		},
	})
}
