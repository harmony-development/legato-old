package socket

import (
	"github.com/sirupsen/logrus"
	"harmony-server/server/http/socket/client"
)

func (h Handler) Deregister(ws *client.Client) {
	guildsQuery, err := h.DB.Query("SELECT guildid FROM guildmembers WHERE userid=$1", ws.UserID)
	if err != nil {
		logrus.Warnf("error deregistering client, potential memory leak : %v", err)
		return
	}
	for guildsQuery.Next() {
		var guildID string
		err = guildsQuery.Scan(&guildID)
		if err != nil {
			logrus.Warnf("Error scanning guilds : %v", err)
			return
		}
		if h.State.Guilds[guildID] != nil && h.State.Guilds[guildID].Clients[*ws.UserID] != nil {
			if len(h.State.Guilds[guildID].Clients[*ws.UserID].Clients) == 1 {
				h.State.Guilds[guildID].Clients[*ws.UserID].Lock()
				h.State.Guilds[guildID].Clients[*ws.UserID] = nil
				h.State.Guilds[guildID].Clients[*ws.UserID].Unlock()
			} else {
				h.State.Guilds[guildID].Clients[*ws.UserID].Lock()
				for i, client := range h.State.Guilds[guildID].Clients[*ws.UserID].Clients {
					if client == ws {
						var c = h.State.Guilds[guildID].Clients[*ws.UserID].Clients
						c[i] = c[len(c)-1]
						h.State.Guilds[guildID].Clients[*ws.UserID].Clients = c[:len(c)-1]
						return
					}
				}
				h.State.Guilds[guildID].Clients[*ws.UserID].Unlock()
			}
		}
	}
}
