package v1

import (
	"github.com/kataras/golog"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket/event"
)

type joinGuildData struct {
	InviteCode string `mapstructure:"invite"`
	Token      string `mapstructure:"token"`
}

func JoinGuild(limiter *rate.Limiter, c echo.Context) error {
	var data joinGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		event.sendErr(ws, "Bad invite code or token")
		return
	}
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
		event.sendErr(ws, "We couldn't get you to join that guild. Please try again")
		golog.Warnf("Error creating joinGuildTransaction : %v", err)
		return
	}
	_, err = joinGuildTransaction.Exec("INSERT INTO guildmembers(userid, guildid) VALUES($1, $2)", ws.Userid, guildid)
	if err != nil {
		event.sendErr(ws, "You are already part of this guild.")
		golog.Warn(err)
		return
	}
	_, err = joinGuildTransaction.Exec("UPDATE invites SET invitecount=invitecount+1 WHERE inviteid=$1", data.InviteCode)
	if err != nil {
		event.sendErr(ws, "We couldn't get you to join that guild. Please try again")
		golog.Warn(err)
		return
	}
	err = joinGuildTransaction.Commit()
	if err != nil {
		event.sendErr(ws, "We couldn't get you to join that guild. Please try again")
		golog.Warnf("Error adding user to guildmembers : %v", err)
		return
	}
	ws.Send(&globals.Packet{
		Type: "joinguild",
		Data: map[string]interface{}{
			"guild": guildid,
		},
	})
	event.registerSocket(guildid, ws, ws.Userid)
}
