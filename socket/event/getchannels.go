package event

import (
	"github.com/kataras/golog"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type getChannelsData struct {
	Token string
	Guild string
}

func OnGetChannels(ws *socket.Client, rawMap map[string]interface{}) {
	var data getChannelsData
	var ok bool
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.Guild, ok = rawMap["guild"].(string); !ok {
		return
	}
	userid := verifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT channelid, channelname FROM channels WHERE guildid=?", data.Guild)
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