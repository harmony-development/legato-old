package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type updateGuildName struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
	Name string `mapstructure:"name"`
}

func OnUpdateGuildName(ws *globals.Client, rawMap map[string]interface{}) {
	var data updateGuildName
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
	_, err := harmonydb.DBInst.Exec("UPDATE guilds SET guildname=$1 WHERE guildid=$2", data.Name, data.Guild)
	if err != nil {
		golog.Warnf("Error updating name. %v", err)
		ws.Send(&globals.Packet{
			Type: "updateguildname",
			Data: map[string]interface{}{
				"success": false,
			},
		})
		return
	}
	for _, client := range globals.Guilds[data.Guild].Clients {
		client.Send(&globals.Packet{
			Type: "updateguildname",
			Data: map[string]interface{}{
				"guild": data.Guild,
				"name": data.Name,
				"success": true,
			},
		})
	}
}