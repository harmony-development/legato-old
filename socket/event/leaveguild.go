package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type leaveGuildData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnLeaveGuild(ws *socket.Client, rawMap map[string]interface{}) {
	var data leaveGuildData
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
	_, err := harmonydb.DBInst.Exec("DELETE FROM guildmembers WHERE userid=? AND guildid=?", userid, data.Guild)
	// GUILD STUCK! GUILD STUCK! PLEASE! I BEG YOU!
	if err != nil {
		golog.Warnf("Error removing member from guild : %v", err)
		ws.Send(&socket.Packet{
			Type: "leaveguild",
			Data: map[string]interface{}{
				"message": "Error leaving guild",
			},
		})
		return
	}
	delete(globals.Guilds[data.Guild].Clients, userid)
	ws.Send(&socket.Packet{
		Type: "leaveguild",
		Data: map[string]interface{}{},
	})
}
