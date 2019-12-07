package event

import (
	"github.com/kataras/golog"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type leaveGuildData struct {
	Token string
	Guild string
}

func OnLeaveGuild(ws *socket.Client, rawMap map[string]interface{}) {
	var data leaveGuildData
	var ok bool
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	userid := verifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	if data.Guild, ok = rawMap["guild"].(string); !ok {
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
	ws.Send(&socket.Packet{
		Type: "leaveguild",
		Data: map[string]interface{}{},
	})
}
