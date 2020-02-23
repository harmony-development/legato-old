package event

import (
	"harmony-server/globals"
	"harmony-server/harmonydb"

	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
)

type getMembersData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnGetMembers(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data getMembersData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		golog.Warnf("Error decoding data while getting members")
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil {
		golog.Warnf("Client tried to list members without being registered")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're getting the guild member list too fast, try again soon")
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT userid FROM guildmembers WHERE guildid=$1", data.Guild)
	if err != nil {
		golog.Warnf("Error getting guild members : %v", err)
		return
	}
	var returnMembers []string
	for res.Next() {
		var userid string
		err = res.Scan(&userid)
		if err != nil {
			golog.Warnf("Error listing userids : %v", err)
			return
		}
		returnMembers = append(returnMembers, userid)
	}

	ws.Send(&globals.Packet{
		Type: "getmembers",
		Data: map[string]interface{}{
			"guild": data.Guild,
			"members": returnMembers,
		},
	})
}
