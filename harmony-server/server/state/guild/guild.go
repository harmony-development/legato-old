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
	Clients map[uint64]*ClientArray
}

// AddClient adds a client to a guild
func (g Guild) AddClient(userID *uint64, c *client.Client) {
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

// Broadcast broadcasts a packet to clients of a guild
func (g Guild) Broadcast(packet *client.OutPacket) {
	for _, client := range g.Clients {
		for _, conn := range client.Clients {
			conn.Send(packet)
		}
	}
}
