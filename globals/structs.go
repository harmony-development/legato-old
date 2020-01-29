package globals

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

type (
	Event func(ws *Client, data map[string]interface{})

	EventBus map[string]Event

	Packet struct {
		Type string `json:"type"`
		Data map[string]interface{} `json:"data"`
	}

	Guild struct {
		Clients map[string][]*Client
		Owner string
	}

	Client struct {
		Connection *websocket.Conn
		EventBus map[string]Event
		Userid string
		LastPong time.Time
		Out chan []byte
	}
)

func (bus EventBus) Bind(s string, event Event) {
	bus[s] = event
}

func (ws *Client) Send(p *Packet) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	ws.Out <- data
}

var Guilds = make(map[string]*Guild)
var Bus EventBus