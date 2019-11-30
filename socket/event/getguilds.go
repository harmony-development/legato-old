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
	Guildid   string `json:"guildid"`
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
	var returnGuilds []guildsData
	for res.Next() {
		var fetchedGuild guildsData
		err := res.Scan(&fetchedGuild.Guildid, &fetchedGuild.Guildname, &fetchedGuild.Picture)
		if err != nil {
			return
		}

		// Now subscribe to all guilds that the client is a member of!
		if globals.Guilds[fetchedGuild.Guildid] != nil {
			globals.Guilds[fetchedGuild.Guildid].Clients[userid] = ws
		} else {
			globals.Guilds[fetchedGuild.Guildid] = &globals.Guild{
				Clients: map[string]*socket.Client{
					userid: ws,
				},
			}
		}
		returnGuilds = append(returnGuilds, fetchedGuild)
	}
	ws.Send(&socket.Packet{
		Type: "getguilds",
		Data: map[string]interface{}{
			"guilds": returnGuilds,
		},
	})
}