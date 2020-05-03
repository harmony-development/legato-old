package guild

import (
	"harmony-server/server/http/socket/handling"
	"sync"
)

// ClientArray is a thread-safe array of client connections
type ClientArray struct {
	*sync.RWMutex
	Clients []*handling.Client // TODO come up with a better name for this
}

// Guild is the data structure for an active guild
type Guild struct {
	*sync.RWMutex
	Clients map[string]*ClientArray
}

func (g Guild) AddClient(userID *string, client *handling.Client) {
	if g.Clients[*userID] == nil {
		g.Lock()
		defer g.Unlock()
		g.Clients[*userID] = &ClientArray{
			Clients: []*handling.Client{client},
		}
		return
	}
	g.Clients[*userID].Lock()
	g.Clients[*userID].Clients = append(g.Clients[*userID].Clients, client)
	g.Clients[*userID].Unlock()
}

func (g Guild) Broadcast(packet *handling.OutPacket) {
	for _, client := range g.Clients {
		for _, conn := range client.Clients {
			conn.Send(packet)
		}
	}
}