package event

import (
	"github.com/kataras/golog"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type joinGuildData struct {
	InviteCode string
	Token string
}

func OnJoinGuild(ws *socket.Client, rawMap map[string]interface{}) {
	var data joinGuildData
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
	if data.InviteCode, ok = rawMap["invitecode"].(string); !ok {
		return
	}
	var guildid string
	err := harmonydb.DBInst.QueryRow("SELECT guildid FROM invites WHERE inviteid=?", data.InviteCode).Scan(&guildid)
	if err != nil {
		golog.Warnf("Error getting invite guild. This probably means something is very screwed up with the guildid. %v", err)
		return
	}
	_, err = harmonydb.DBInst.Exec("INSERT INTO guildmembers(userid, guildid) VALUES(?, ?)", userid, guildid)
	if err != nil {
		golog.Warnf("Error adding user to guildmembers : %v", err)
		return
	}
	registerSocket(guildid, ws, userid)
	ws.Send(&socket.Packet{
		Type: "joinguild",
		Data: map[string]interface{}{
			"guild": guildid,
		},
	})
}
