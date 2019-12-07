package event

import (
	"github.com/kataras/golog"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type updateGuildPictureData struct {
	Token string
	Guild string
	Picture string
}

func OnUpdateGuildPicture(ws *socket.Client, rawMap map[string]interface{}) {
	var ok bool
	var data updateGuildPictureData
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.Guild, ok = rawMap["guild"].(string); !ok {
		return
	}
	if data.Picture, ok = rawMap["picture"].(string); !ok {
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
	_, err := harmonydb.DBInst.Exec("UPDATE guilds SET picture=? WHERE guildid=?", data.Picture, data.Guild)
	if err != nil {
		golog.Warnf("Error updating picture. %v", err)
		ws.Send(&socket.Packet{
			Type: "updateguildpicture",
			Data: map[string]interface{}{
				"success": false,
			},
		})
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		client.Send(&socket.Packet{
			Type: "updateguildpicture",
			Data: map[string]interface{}{
				"guild": data.Guild,
				"picture": data.Picture,
				"success": true,
			},
		})
	}
}