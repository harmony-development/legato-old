package socket

import (
	"harmony-server/server/http/socket/client"

	"github.com/sirupsen/logrus"
)

// Deregister terminates a client's session
func (h Handler) Deregister(ws *client.Client) {
	if ws.UserID == nil {
		return
	}
	guilds, err := h.DB.GuildsForUser(*ws.UserID)
	if err != nil {
		logrus.Warnf("error deregistering client, potential memory leak : %v", err)
		return
	}
	for _, guildID := range guilds {
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
