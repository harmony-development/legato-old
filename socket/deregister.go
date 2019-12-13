package socket

import (
	"github.com/kataras/golog"
	"harmony-server/globals"
	"harmony-server/harmonydb"
)

func deregister(ws *Client) {
	guildsQuery, err := harmonydb.DBInst.Query("SELECT guildid FROM guildmembers WHERE userid=$1", ws.Userid)
	if err != nil {
		golog.Warnf("ERROR deregistering socket! POTENTIAL MEMORY LEAK! %v", err)
		return
	}
	for guildsQuery.Next() {
		var guildID string
		err = guildsQuery.Scan(&guildID)
		if err != nil {
			golog.Warnf("Error scanning guilds : %v", err)
			return
		}
		delete(globals.Guilds[guildID].Clients, ws.Userid)
	}
}
