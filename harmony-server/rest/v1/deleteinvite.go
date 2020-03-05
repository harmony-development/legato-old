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

type deleteInviteData struct {
	Token  string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
	Invite string `mapstructure:"invite"`
}

func DeleteInvite(limiter *rate.Limiter, ctx echo.Context) error {
	var data deleteInviteData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil || globals.Guilds[data.Guild].Owner != ws.Userid {
		return
	}
	if !limiter.Allow() {
		event.sendErr(ws, "You're deleting a lot of invites, wait a sec and try again")
		return
	}
	_, err := harmonydb.DBInst.Exec("DELETE FROM invites WHERE inviteid=$1 AND guildid=$2", data.Invite, data.Guild)
	if err != nil {
		event.sendErr(ws, "We weren't able to delete that invite for some reason. You should try again")
		golog.Warnf("Error deleting invite : %v", err)
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
