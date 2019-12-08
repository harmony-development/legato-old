package event

import (
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type deleteInviteData struct {
	Token  string
	Guild string
	Invite string
}

func OnDeleteInvite(ws *socket.Client, rawMap map[string]interface{}) {
	var data deleteInviteData
	var ok bool
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	if data.Invite, ok = rawMap["invite"].(string); !ok {
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
	_, err := harmonydb.DBInst.Exec("DELETE FROM invites WHERE inviteid=? AND guildid=?", data.Invite, data.Guild)
	if err != nil {
		return
	}
	ws.Send(&socket.Packet{
		Type: "deleteinvite",
		Data: map[string]interface{}{
			"success": true,
			"invite": data.Invite,
		},
	})
}
