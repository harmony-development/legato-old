package socket

import (
	"encoding/json"
	"github.com/gorilla/websocket"

	"harmony-server/globals"
	"harmony-server/socket/event"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:    2048,
		WriteBufferSize:   2048,
		CheckOrigin:       func(r *http.Request) bool { // we will allow all domains... For now...
			return true
		},
		EnableCompression: true,
	}
)

func NewSocket(w http.ResponseWriter, r *http.Request) *globals.Client {
	rawsocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Warnf("error upgrading event for reason : %v", err)
	}
	ws := &globals.Client{
		Connection: rawsocket,
		EventBus: make(map[string]globals.Event),
		Out: make(chan []byte),
		Authed: false,
	}
	go reader(ws)
	go writer(ws)
	go pinger(ws)
	return ws
}

func pinger(ws *globals.Client) {
	ws.Send(&globals.Packet{
		Type: "ping",
		Data: nil,
	})
	time.Sleep(20 * time.Second)
	if time.Since(ws.LastPong) > 20 * time.Second {
		deregister(ws)
		logrus.Debugf("Closing Socket : Ping Timeout")
		err := ws.Connection.Close()
		if err != nil {
			logrus.Warnf("Error closing websocket connection : %v", err)
		}
	}
}

// reader eternally waits for things to read from the client
func reader(ws *globals.Client) {
	for {
		_, msg, err := ws.Connection.ReadMessage()
		if err == nil {
			var p globals.Packet
			if err = json.Unmarshal(msg, &p); err == nil {
				if ws.EventBus[p.Type] != nil {
					// login/register/auth/ping do not require authentication, everything else does
					if ws.Authed || p.Type == "login" || p.Type == "register" || p.Type == "auth" || p.Type == "ping" {
						ws.EventBus[p.Type](ws, p.Data) // call an event from the eventbus if it exists
					} else {
						event.Deauth(ws)
					}
				} else {
					logrus.Warnf("Unrecognized API Query Detected : %v", p.Type)
				}
			}
		} else {
			logrus.Warnf("Error reading data from event : %v", err)
			logrus.Debugf("Closing Socket : Data read error")
			if ws.Userid != "" {
				deregister(ws)
			}
			_ = ws.Connection.Close()
			return
		}
	}
}

// writer eternally waits for things to write to the event
func writer(ws *globals.Client) {
	for {
		msg := <- ws.Out // wait for a new message to be sent
		err := ws.Connection.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			logrus.Warnf("Error writing data to event : %v", err)
			logrus.Debugf("Closing Socket : Data write error")
			_ = ws.Connection.Close()
			if ws.Userid != "" {
				deregister(ws)
			}
			return
		}
	}
}