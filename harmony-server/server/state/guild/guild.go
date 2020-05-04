package guild

import (
	"harmony-server/server/http/socket/client"
	"sync"
)

// ClientArray is a thread-safe array of client connections
type ClientArray struct {
	*sync.RWMutex
	Clients []*client.Client // TODO come up with a better name for this
}

// Guild is the data structure for an active guild
type Guild struct {
	*sync.RWMutex
	Clients map[string]*ClientArray
}

func (g Guild) AddClient(userID *string, c *client.Client) {
	if g.Clients[*userID] == nil {
		g.Lock()
		defer g.Unlock()
		g.Clients[*userID] = &ClientArray{
			Clients: []*client.Client{},
		}
		return
	}
	g.Clients[*userID].Lock()
	g.Clients[*userID].Clients = append(g.Clients[*userID].Clients, c)
	g.Clients[*userID].Unlock()
}

func (g Guild) Broadcast(packet *client.OutPacket) {
	for _, client := range g.Clients {
		for _, conn := range client.Clients {
			conn.Send(packet)
		}
	}
}
