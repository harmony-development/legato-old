package event

import (
	"github.com/kataras/golog"
	"github.com/thanhpk/randstr"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type addChannelData struct {
	Token       string
	Guild       string
	Channelname string
}

func OnAddChannel(ws *socket.Client, rawMap map[string]interface{}) {
	var data addChannelData
	var ok bool
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.Guild, ok = rawMap["guild"].(string); !ok {
		return
	}
	if data.Channelname, ok = rawMap["channelname"].(string); !ok {
		return
	}
	userid := verifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil || globals.Guilds[data.Guild].Owner != userid {
		return
	}
	var channelID = randstr.Hex(16)
	_, err := harmonydb.DBInst.Exec("INSERT INTO channels(channelid, guildid, channelname) VALUES(?, ?, ?)", channelID, data.Guild, data.Channelname)
	if err != nil {
		golog.Warnf("Error creating channel : %v", err)
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		client.Send(&socket.Packet{
			Type: "addguildchannel",
			Data: map[string]interface{}{
				"guild":   data.Guild,
				"channelname": data.Channelname,
				"channelid": channelID,
				"success": true,
			},
		})
	}
}
