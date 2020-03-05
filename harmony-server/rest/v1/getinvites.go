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

type getInvitesData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func GetInvites(limiter *rate.Limiter, ctx echo.Context) error {
	var data getInvitesData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil || globals.Guilds[data.Guild].Owner != ws.Userid {
		return
	}
	if !limiter.Allow() {
		event.sendErr(ws, "You're listing invites too fast, slow down for a bit and try again")
		return
	}
	res, err := harmonydb.DBInst.Query("SElECT inviteid, invitecount FROM invites WHERE guildid=$1 ORDER BY invitecount", data.Guild)
	if err != nil {
		event.sendErr(ws, "We weren't able to get a list of invites for this guild. Please try again")
		golog.Warnf("Error getting invites : %v", err)
		return
	}
	returnInvites := make(map[string]int)
	for res.Next() {
		var invitecode string
		var invitecount int
		err = res.Scan(&invitecode, &invitecount)
		if err != nil {
			event.sendErr(ws, "We weren't able to get a list of invites for this guild. Please try again")
			golog.Warnf("Error scanning invite codes : %v", err)
			return
		}
		returnInvites[invitecode] = invitecount
	}
	ws.Send(&globals.Packet{
		Type: "getinvites",
		Data: map[string]interface{}{
			"invites": returnInvites,
			"guild":   data.Guild,
		},
	})
}
