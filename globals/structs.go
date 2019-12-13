package globals

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

type (
	Event func(ws *Client, data map[string]interface{})

	Packet struct {
		Type string `json:"type"`
		Data map[string]interface{} `json:"data"`
	}

	Guild struct {
		Clients map[string]*Client
		Owner string
	}

	Client struct {
		Connection *websocket.Conn
		EventBus map[string]Event
		Userid string
		Out chan []byte
	}
)

func (ws *Client) Bind(s string, event Event) {
	ws.EventBus[s] = event
}

func (ws *Client) Send(p *Packet) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	ws.Out <- data
}

var Guilds = make(map[string]*Guild)