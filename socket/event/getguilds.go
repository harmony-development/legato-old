package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"harmony-server/globals"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type getGuildsData struct {
	Token string `mapstructure:"token"`
}

type guildsData struct {
	Guildname string `json:"guildname"`
	Picture   string `json:"picture"`
	IsOwner   bool   `json:"owner"`
}

func OnGetGuilds(ws *socket.Client, rawMap map[string]interface{}) {
	var data getGuildsData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	res, err := harmonydb.DBInst.Query("SELECT guilds.guildid, guilds.guildname, guilds.owner, guilds.picture FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=?", userid)
	if err != nil {
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
			golog.Warnf("Error scanning next row. Reason: %v", err)
			return
		}
		// Now subscribe to all guilds that the client is a member of!
		if globals.Guilds[guildID] != nil {
			globals.Guilds[guildID].Clients[userid] = ws
			globals.Guilds[guildID].Owner = guildOwner
		} else {
			globals.Guilds[guildID] = &globals.Guild{
				Clients: map[string]*socket.Client{
					userid: ws,
				},
			}
			globals.Guilds[guildID].Owner = guildOwner
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
