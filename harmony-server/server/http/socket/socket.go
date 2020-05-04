package socket

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"harmony-server/server/db"
	"harmony-server/server/http/socket/client"
	"harmony-server/server/state"
	"net/http"
)



// Handler is an instance of the socket handler
type Handler struct {
	Upgrader *websocket.Upgrader
	DB       *db.DB
	Bus      client.Bus
	State    *state.State
}

// NewHandler creates a new socket handler
func NewHandler(state *state.State) *Handler {
	var bus = make(client.Bus)
	h := &Handler{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  2048,
			WriteBufferSize: 2048,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			EnableCompression: true,
		},
		Bus: bus,
		State: state,
	}
	h.Setup()
	return h
}

// Handle upgrades an HTTP request to a Client
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) *client.Client {
	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Warnf("error upgrading events", err)
		return nil
	}

	c := &client.Client{
		Conn: conn,
		Bus:  h.Bus,
		Out:  make(chan []byte),
	}

	go c.Reader()
	go c.Writer()
	go c.Pinger()

	return c
}


