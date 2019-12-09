package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type getInvitesData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnGetInvites(ws *socket.Client, rawMap map[string]interface{}) {
	var data getInvitesData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	if userid == "" { // token is invalid! Get outta here!
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil || globals.Guilds[data.Guild].Owner != userid {
		return
	}
	res, err := harmonydb.DBInst.Query("SElECT inviteid, invitecount FROM invites WHERE guildid=$1 ORDER BY invitecount", data.Guild)
	if err != nil {
		golog.Warnf("Error getting invites : %v", err)
		return
	}
	returnInvites := make(map[string]int)
	for res.Next() {
		var invitecode string
		var invitecount int
		err = res.Scan(&invitecode, &invitecount)
		if err != nil {
			golog.Warnf("Error scanning invite codes : %v", err)
			return
		}
		returnInvites[invitecode] = invitecount
	}
	ws.Send(&socket.Packet{
		Type: "getinvites",
		Data: map[string]interface{}{
			"invites": returnInvites,
			"guild":   data.Guild,
		},
	})
}
