package socket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	. "github.com/logrusorgru/aurora"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket is a structure that contains all low-level server data
type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]EventHandler
}

// CreateServer creates a websocket server that handles the clients
func CreateServer(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(Red("Error upgrading : " + err.Error()).Bold())
		return nil, err
	}

	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte),
		In:     make(chan []byte),
		Events: make(map[string]EventHandler),
	}

	go ws.Reader()
	go ws.Writer()
	return ws, nil
}

// Reader handles reading for the socket server
func (ws *WebSocket) Reader() {
	defer func() {
		ws.Conn.Close()
	}()
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			break
		}
		event, err := ParseMessage(message)
		if err != nil {
			log.Println("Unable to parse the receiver's data")
			break
		}
		if handler, exists := ws.Events[event.Name]; exists {
			handler(event.Data, ws)
		}
	}
}

// Writer handles writing for the socket server
func (ws *WebSocket) Writer() {
	for {
		select {
		case message, exists := <-ws.Out:
			if !exists {
				ws.Conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
				return
			}
			w, err := ws.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			w.Close()
		}
	}
}

// On is a helper function that registers an event handler for a specific event
func (ws *WebSocket) On(event string, callback func(data interface{}, ws *WebSocket)) {
	ws.Events[event] = callback
}
