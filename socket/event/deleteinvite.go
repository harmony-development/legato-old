package event

import (
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type deleteInviteData struct {
	Token  string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
	Invite string `mapstructure:"invite"`
}

func OnDeleteInvite(ws *globals.Client, rawMap map[string]interface{}) {
	var data deleteInviteData
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
	_, err := harmonydb.DBInst.Exec("DELETE FROM invites WHERE inviteid=$1 AND guildid=$2", data.Invite, data.Guild)
	if err != nil {
		return
	}
	ws.Send(&globals.Packet{
		Type: "deleteinvite",
		Data: map[string]interface{}{
			"success": true,
			"invite": data.Invite,
		},
	})
}
