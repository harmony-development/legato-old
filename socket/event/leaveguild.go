package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type leaveGuildData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnLeaveGuild(ws *globals.Client, rawMap map[string]interface{}) {
	var data leaveGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid ,err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[userid] == nil {
		return
	}
	_, err = harmonydb.DBInst.Exec("DELETE FROM guildmembers WHERE userid=$1 AND guildid=$2", userid, data.Guild)
	// GUILD STUCK! GUILD STUCK! PLEASE! I BEG YOU!
	if err != nil {
		golog.Warnf("Error removing member from guild : %v", err)
		sendErr(ws, "Uh oh. Seems like you couldn't leave the guild. Please try again. If this keeps happening, please contact administration")
		return
	}
	delete(globals.Guilds[data.Guild].Clients, userid)
	ws.Send(&globals.Packet{
		Type: "leaveguild",
		Data: map[string]interface{}{},
	})
}
