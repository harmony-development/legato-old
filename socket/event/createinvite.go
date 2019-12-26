package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"github.com/thanhpk/randstr"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type createInviteData struct {
	Token  string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnCreateInvite(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data createInviteData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid, err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil || globals.Guilds[data.Guild].Owner != userid {
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "That's quite a lot of invites... Try again in a few secs")
		return
	}
	var inviteID = randstr.Hex(5)
	_, err = harmonydb.DBInst.Exec("INSERT INTO invites(inviteid, guildid) VALUES($1, $2)", inviteID, data.Guild)
	if err != nil {
		sendErr(ws, "We weren't able to make an invite link. Please try again")
		golog.Warnf("Error inserting invite : %v", err)
		return
	}
	ws.Send(&globals.Packet{
		Type: "createinvite",
		Data: map[string]interface{}{
			"success": true,
			"invite": inviteID,
		},
	})
}