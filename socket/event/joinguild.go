package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type joinGuildData struct {
	InviteCode string `mapstructure:"invite"`
	Token      string `mapstructure:"token"`
}

func OnJoinGuild(ws *socket.Client, rawMap map[string]interface{}) {
	var data joinGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	var guildid string
	err := harmonydb.DBInst.QueryRow("SELECT guildid FROM invites WHERE inviteid=?", data.InviteCode).Scan(&guildid)
	if err != nil {
		ws.Send(&socket.Packet{
			Type: "joinguild",
			Data: map[string]interface{}{
				"message": "Invalid Invite Code!",
			},
		})
		golog.Warnf("Error getting invite guild. This probably means the guild invite code doesn't exist. %v", err)
		return
	}
	_, err = harmonydb.DBInst.Exec("INSERT INTO guildmembers(userid, guildid) VALUES(?, ?); UPDATE invites SET invitecount=invitecount+1 WHERE inviteid=?", userid, guildid, data.InviteCode)
	if err != nil {
		ws.Send(&socket.Packet{
			Type: "joinguild",
			Data: map[string]interface{}{
				"message": "Error Joining Guild!",
			},
		})
		golog.Warnf("Error adding user to guildmembers : %v", err)
		return
	}
	ws.Send(&socket.Packet{
		Type: "joinguild",
		Data: map[string]interface{}{
			"guild": guildid,
		},
	})
	registerSocket(guildid, ws, userid)
}
