package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type subscribeToGuildData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnSubscribeToGuild(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data subscribeToGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid, err := authentication.VerifyToken(data.Token)
	if err != nil {
		sendErr(ws, "invalid token")
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "Too many subscription attempts, please try later")
		return
	}
	var count int
	res, err := harmonydb.DBInst.Query("SELECT COUNT(*) FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1 AND guilds.guildid=$2", userid, data.Guild)
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
		RegisterSocket(data.Guild, ws, userid)
	}
	return
}