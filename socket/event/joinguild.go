package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type joinGuildData struct {
	InviteCode string `mapstructure:"invite"`
	Token      string `mapstructure:"token"`
}

func OnJoinGuild(ws *globals.Client, rawMap map[string]interface{}) {
	var data joinGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	var guildid string
	err := harmonydb.DBInst.QueryRow("SELECT guildid FROM invites WHERE inviteid=$1", data.InviteCode).Scan(&guildid)
	if err != nil {
		ws.Send(&globals.Packet{
			Type: "joinguild",
			Data: map[string]interface{}{
				"message": "Invalid Invite Code!",
			},
		})
		golog.Warnf("Error getting invite guild. This probably means the guild invite code doesn't exist. %v", err)
		return
	}
	joinGuildTransaction, err := harmonydb.DBInst.Begin()
	if err != nil {
		golog.Warnf("Error creating joinGuildTransaction : %v", err)
		return
	}
	_, err = joinGuildTransaction.Exec("INSERT INTO guildmembers(userid, guildid) VALUES($1, $2)", userid, guildid)
	if err != nil {
		golog.Warn(err)
		return
	}
	_, err = joinGuildTransaction.Exec("UPDATE invites SET invitecount=invitecount+1 WHERE inviteid=$1", data.InviteCode)
	if err != nil {
		golog.Warn(err)
		return
	}
	err = joinGuildTransaction.Commit()
	if err != nil {
		ws.Send(&globals.Packet{
			Type: "joinguild",
			Data: map[string]interface{}{
				"message": "Error Joining Guild!",
			},
		})
		golog.Warnf("Error adding user to guildmembers : %v", err)
		return
	}
	ws.Send(&globals.Packet{
		Type: "joinguild",
		Data: map[string]interface{}{
			"guild": guildid,
		},
	})
	registerSocket(guildid, ws, userid)
}
