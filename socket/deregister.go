package socket

import (
	"github.com/kataras/golog"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

func deregister(ws *globals.Client) {
	guildsQuery, err := harmonydb.DBInst.Query("SELECT guildid FROM guildmembers WHERE userid=$1", ws.Userid)
	if err != nil {
		golog.Warnf("error deregistering client, potential memory leak : %v", err)
		return
	}
	for guildsQuery.Next() {
		var guildID string
		err = guildsQuery.Scan(&guildID)
		if err != nil {
			golog.Warnf("Error scanning guilds : %v", err)
			return
		}
		if globals.Guilds[guildID].Clients[ws.Userid] != nil {
			for i, client := range globals.Guilds[guildID].Clients[ws.Userid] {
				if client == ws {
					globals.Guilds[guildID].Clients[ws.Userid] = append(globals.Guilds[guildID].Clients[ws.Userid][:i], globals.Guilds[guildID].Clients[ws.Userid][i:]...)
				}
			}
		}
	}
}
