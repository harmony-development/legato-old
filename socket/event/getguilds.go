package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/time/rate"
	"harmony-server/authentication"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

type getGuildsData struct {
	Token string `mapstructure:"token"`
}

type guildsData struct {
	Guildname string `json:"guildname"`
	Picture   string `json:"picture"`
	IsOwner   bool   `json:"owner"`
}

func OnGetGuilds(ws *globals.Client, rawMap map[string]interface{}, limiter *rate.Limiter) {
	var data getGuildsData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid ,err := authentication.VerifyToken(data.Token)
	if err != nil {
		deauth(ws)
		return
	}
	if !limiter.Allow() {
		sendErr(ws, "You're getting the guilds list a lot, please try again in a sec")
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT guilds.guildid, guilds.guildname, guilds.owner, guilds.picture FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=$1", userid)
	if err != nil {
		sendErr(ws, "We weren't able to get a list of guilds for you. Try reloading the page / logging back in")
		golog.Warnf("Error selecting guilds. Reason : %v", err)
		return
	}
	var returnGuilds = make(map[string]guildsData)
	for res.Next() {
		var guildID string
		var fetchedGuild guildsData
		var guildOwner string
		err := res.Scan(&guildID, &fetchedGuild.Guildname, &guildOwner, &fetchedGuild.Picture)
		if guildOwner == userid {
			fetchedGuild.IsOwner = true
		}
		if err != nil {
			sendErr(ws, "We weren't able to get a list of guilds for you. Try reloading the page / logging back in")
			golog.Warnf("Error scanning next row. Reason: %v", err)
			return
		}
		// Now subscribe to all guilds that the client is a member of!
		registerSocket(guildID, ws, userid)
		returnGuilds[guildID] = fetchedGuild
	}
	ws.Send(&globals.Packet{
		Type: "getguilds",
		Data: map[string]interface{}{
			"guilds": returnGuilds,
		},
	})
}
