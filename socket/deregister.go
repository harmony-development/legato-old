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
			if len(globals.Guilds[guildID].Clients[ws.Userid]) == 1 {
				globals.Guilds[guildID].Clients[ws.Userid] = nil
				return
			}
			for i, client := range globals.Guilds[guildID].Clients[ws.Userid] {
				if client == ws {
					var c = globals.Guilds[guildID].Clients[ws.Userid]
					c[i] = c[len(c)-1]
					globals.Guilds[guildID].Clients[ws.Userid] = c[:len(c)-1]
					return
				}
			}
		}
	}
}
