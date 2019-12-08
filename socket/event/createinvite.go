package event

import (
	"github.com/thanhpk/randstr"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type createInviteData struct {
	Token  string
	Guild string
}

func OnCreateInvite(ws *socket.Client, rawMap map[string]interface{}) {
	var data createInviteData
	var ok bool
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.Guild, ok = rawMap["guild"].(string); !ok {
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
	var inviteID = randstr.Hex(5)
	_, err := harmonydb.DBInst.Exec("INSERT INTO invites(inviteid, guildid) VALUES(?, ?)", inviteID, data.Guild)
	if err != nil {
		return
	}
	ws.Send(&socket.Packet{
		Type: "createinvite",
		Data: map[string]interface{}{
			"success": true,
			"invite": inviteID,
		},
	})
}