package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/db"
)

type subscribeToGuildData struct {
	Session string `mapstructure:"session"`
	Guild string `mapstructure:"guild"`
}

// OnSubscribeToGuild handles requests to subscribe to a guilds events
func OnSubscribeToGuild(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data subscribeToGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid, err := authentication.GetUserBySession(data.Session)
	if err != nil {
		sendErr(ws, "invalid token")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "Too many subscription attempts, please try later")
		return
	}
	var count int
	res, err := db.DBInst.Query("SELECT COUNT(*) FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1 AND guilds.guildid=$2", userid, data.Guild)
	if err != nil {
		sendErr(ws, "unable to subscribe to guild, try reloading?")
		return
	}
	err = res.Scan(&count)
	if err != nil {
		sendErr(ws,  "unable to subscribe to guild, try reloading?")
		return
	}
	if count == 1 {
		RegisterSocket(data.Guild, ws, data.Session)
	}
	return
}