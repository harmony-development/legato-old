package event

import (
	"github.com/kataras/golog"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type getServersData struct {
	Token string
}

type serverData struct {
	Guildid string `json:"guildid"`
	Servername string `json:"servername"`
	Picture string `json:"picture"`
}

func OnGetServers(ws *socket.Client, rawMap map[string]interface{}) {
	var data getServersData
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
	res, err := harmonydb.DBInst.Query("SELECT guilds.guildid, guilds.servername, guilds.picture FROM guildmembers INNER JOIN guilds ON guildmembers.guildid = guilds.guildid WHERE userid=?", userid)
	if err != nil {
		golog.Warnf("Error selecting guilds. Reason : %v", err)
		return
	}
	var returnServers []serverData
	for res.Next() {
		var fetchedServer serverData
		err := res.Scan(&fetchedServer.Guildid, &fetchedServer.Servername, &fetchedServer.Picture)
		if err != nil {
			return
		}
		returnServers = append(returnServers, fetchedServer)
	}
	ws.Send(&socket.Packet{
		Type: "GetServers",
		Data: map[string]interface{}{
			"servers": returnServers,
		},
	})
}