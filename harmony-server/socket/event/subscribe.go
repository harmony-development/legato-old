package event

import (
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/db"
	"harmony-server/globals"
)

type subscribeData struct {
	Session string `mapstructure:"session"`
}

func OnSubscribe(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data subscribeData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	user, err := authentication.GetUserBySession(data.Session)
	if err != nil {
		sendErr(ws, "invalid session")
		return
	}
	ws.Userid = user.ID
	if !limiter.Allow() {
		sendErr(ws, "Too many subscription attempts, please try later")
		return
	}
	res, err := db.DBInst.Query("SELECT guilds.guildid FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1", user.ID)
	if err != nil {
		sendErr(ws, "We weren't able to get a list of guilds. Try reloading the page / logging back in")
		logrus.Warnf("Error selecting guilds. Reason : %v", err)
		return
	}
	for res.Next() {
		var guildID string
		err := res.Scan(&guildID)
		if err != nil {
			sendErr(ws, "Unable to subscribe to all guilds. Please try again later")
			logrus.Warnf("Error scanning next row. Reason: %v", err)
			return
		}
		// Now subscribe to all guilds that the client is a member of!
		RegisterSocket(guildID, ws, ws.Userid)
	}
}
