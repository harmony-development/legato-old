package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
	"path"
)

type updateGuildPictureData struct {
	Token   string `mapstructure:"token"`
	Guild   string `mapstructure:"guild"`
	Picture string `mapstructure:"picture"`
}

func OnUpdateGuildPicture(ws *socket.Client, rawMap map[string]interface{}) {
	var data updateGuildPictureData
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
	var oldPictureID string
	err := harmonydb.DBInst.QueryRow("SELECT picture FROM guilds WHERE guildid=$1", data.Guild).Scan(&oldPictureID)
	_, err = harmonydb.DBInst.Exec("UPDATE guilds SET picture=$1 WHERE guildid=$2", data.Picture, data.Guild)
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
	go DeleteFromFilestore(path.Base(oldPictureID))
	for _, client := range globals.Guilds[data.Guild].Clients {
		client.Send(&socket.Packet{
			Type: "updateguildpicture",
			Data: map[string]interface{}{
				"guild":   data.Guild,
				"picture": data.Picture,
				"success": true,
			},
		})
	}
}
