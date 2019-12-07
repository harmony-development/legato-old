package socket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/kataras/golog"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:    2048,
		WriteBufferSize:   2048,
		CheckOrigin:       func(r *http.Request) bool { // we will allow all domains... For now...
			return true
		},
		EnableCompression: false,
	}
)

type (
	Event func(ws *Client, data map[string]interface{})

	Packet struct {
		Type string `json:"type"`
		Data map[string]interface{} `json:"data"`
	}

	Client struct {
		Connection *websocket.Conn
		EventBus map[string]Event
		Out chan []byte
	}
)

func NewSocket(w http.ResponseWriter, r *http.Request) *Client {
	rawsocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		golog.Warnf("error upgrading event for reason : %v", err)
	}
	ws := &Client{
		Connection: rawsocket,
		EventBus: make(map[string]Event),
		Out: make(chan []byte),
	}
	go reader(ws)
	go writer(ws)
	return ws
}

// reader eternally waits for things to read from the event
func reader(ws *Client) {
	for {
		_, msg, err := ws.Connection.ReadMessage()
		if err == nil {
			var p Packet
			if err = json.Unmarshal(msg, &p); err == nil {
				if ws.EventBus[p.Type] != nil {
					ws.EventBus[p.Type](ws, p.Data) // call an event from the eventbus if it exists
				} else {
					golog.Warnf("Unrecognized API Query Detected : %v", p.Type)
				}
			}
		} else {
			golog.Warnf("Error reading data from event : %v", err)
			return
		}
	}
}

// writer eternally waits for things to write to the event
func writer(ws *Client) {
	for {
		msg := <- ws.Out // wait for a new message to be sent
		err := ws.Connection.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			golog.Warnf("Error writing data to event : %v", err)
		}
	}
}

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