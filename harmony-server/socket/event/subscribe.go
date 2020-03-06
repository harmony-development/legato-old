package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type subscribeData struct {
	Token string `mapstructure:"token"`
}

func OnSubscribe(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data subscribeData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid, err := authentication.VerifyToken(data.Token)
	if err != nil {
		sendErr(ws, "invalid token")
		return
	}
	ws.Userid = userid
	if !limiter.Allow() {
		sendErr(ws, "Too many subscription attempts, please try later")
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT guilds.guildid FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1", userid)
	if err != nil {
		sendErr(ws, "We weren't able to get a list of guilds. Try reloading the page / logging back in")
		golog.Warnf("Error selecting guilds. Reason : %v", err)
		return
	}
	for res.Next() {
		var guildID string
		err := res.Scan(&guildID)
		if err != nil {
			sendErr(ws, "Unable to subscribe to all guilds. Please try again later")
			golog.Warnf("Error scanning next row. Reason: %v", err)
			return
		}
		// Now subscribe to all guilds that the client is a member of!
		RegisterSocket(guildID, ws, ws.Userid)
	}
}
