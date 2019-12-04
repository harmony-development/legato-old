package event

import (
	"github.com/kataras/golog"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type getGuildsData struct {
	Token string
}

type guildsData struct {
	Guildname string `json:"guildname"`
	Picture   string `json:"picture"`
}

func OnGetGuilds(ws *socket.Client, rawMap map[string]interface{}) {
	var data getGuildsData
	var ok bool
	if data.Token, ok = rawMap["token"].(string); !ok {
		deauth(ws)
		return
	}
	userid := verifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT guilds.guildid, guilds.guildname, guilds.picture FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=?", userid)
	if err != nil {
		golog.Warnf("Error selecting guilds. Reason : %v", err)
		return
	}
	var returnGuilds = make(map[string]guildsData)
	for res.Next() {
		var guildID string
		var fetchedGuild guildsData
		err := res.Scan(&guildID, &fetchedGuild.Guildname, &fetchedGuild.Picture)
		if err != nil {
			golog.Warnf("Error scanning next row. Reason: %v", err)
			return
		}
		// Now subscribe to all guilds that the client is a member of!
		if globals.Guilds[guildID] != nil {
			globals.Guilds[guildID].Clients[userid] = ws
		} else {
			globals.Guilds[guildID] = &globals.Guild{
				Clients: map[string]*socket.Client{
					userid: ws,
				},
			}
		}
		returnGuilds[guildID] = fetchedGuild
	}
	ws.Send(&socket.Packet{
		Type: "getguilds",
		Data: map[string]interface{}{
			"guilds": returnGuilds,
		},
	})
}