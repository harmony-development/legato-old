package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type subscribeData struct {
	Token string `mapstructure:"token"`
}

type guildsData struct {
	Guildname string `json:"guildname"`
	Picture   string `json:"picture"`
	IsOwner   bool   `json:"owner"`
}

func OnSubscribe(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data subscribeData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're subscribing too often, please try again later")
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT guilds.guildid FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1", ws.Userid)
	if err != nil {
		sendErr(ws, "We weren't able to get a list of guilds for you. Try reloading the page / logging back in")
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
		registerSocket(guildID, ws, ws.Userid)
	}
}
