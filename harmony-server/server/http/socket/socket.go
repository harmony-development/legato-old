package socket

import (
	"harmony-server/server/db"
	"harmony-server/server/http/socket/client"
	"harmony-server/server/logger"
	"harmony-server/server/state"
	"net/http"

	"github.com/gorilla/websocket"
)

// Handler is an instance of the socket handler
type Handler struct {
	Upgrader *websocket.Upgrader
	DB       *db.DB
	Logger   *logger.Logger
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
		Bus:   bus,
		State: state,
	}
	h.Setup()
	return h
}

// Handle upgrades an HTTP request to a Client
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) *client.Client {
	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Logger.Exception(err)
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
