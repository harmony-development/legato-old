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

type leaveGuildData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func LeaveGuild(limiter *rate.Limiter, c echo.Context) error {
	var data leaveGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if globals.Guilds[data.Guild] == nil || globals.Guilds[data.Guild].Clients[ws.Userid] == nil {
		return
	}
	var ownerID  string
	err := harmonydb.DBInst.QueryRow("SELECT owner FROM guilds WHERE guildid=$1", data.Guild).Scan(&ownerID)
	if err != nil {
		golog.Warnf("Error leaving guild : %v", err)
		event.sendErr(ws, "We were unable to get you to leave the guild for some reason")
		return
	}
	if ownerID == ws.Userid {
		event.sendErr(ws, "You cannot leave a guild you own")
		return
	}
	_, err = harmonydb.DBInst.Exec("DELETE FROM guildmembers WHERE userid=$1 AND guildid=$2", ws.Userid, data.Guild)
	// GUILD STUCK! GUILD STUCK! PLEASE! I BEG YOU!
	if err != nil {
		golog.Warnf("Error removing member from guild : %v", err)
		event.sendErr(ws, "Uh oh. Seems like you couldn't leave the guild. Please try again. If this keeps happening, please contact administration")
		return
	}
	delete(globals.Guilds[data.Guild].Clients, ws.Userid)
	ws.Send(&globals.Packet{
		Type: "leaveguild",
		Data: map[string]interface{}{},
	})
}
