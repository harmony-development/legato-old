package socket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"harmony-server/server/db"
	"harmony-server/server/state"
	"harmony-server/server/state/event"
	"harmony-server/server/state/guild"
	"net/http"
	"sync"
	"time"
)

// Client is the data structure for a connected client
type Client struct {
	*sync.RWMutex
	Conn     *websocket.Conn
	Bus      event.Bus
	UserID   *string
	LastPong time.Time
	Out      chan []byte
}

// Handler is an instance of the socket handler
type Handler struct {
	Upgrader *websocket.Upgrader
	DB       *db.DB
	Bus      event.Bus
	State    *state.State
}

// Packet is the standard way all messages are delivered and received from the server
type Packet struct {
	Type string
	Data *json.RawMessage
}

// OutPacket is like packet but the data doesn't necessarily have to be raw bytes
type OutPacket struct {
	Type string
	Data interface{}
}

// NewHandler creates a new socket handler
func NewHandler(state *state.State) *Handler {
	var bus = make(event.Bus)
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
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) *guild.Client {
	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Warnf("error upgrading events", err)
		return nil
	}

	c := &Client{
		Conn: conn,
		Bus:  h.Bus,
		Out:  make(chan []byte),
	}

	go c.Reader()
	go c.Writer()
	go c.Pinger()

	return c
}

// Send adds a packet to the socket queue
func (c *Client) Send(p *OutPacket) {
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	c.Out <- data
}

// SendError adds a packet to the socket queue
func (c *Client) SendError(msg string) {
	c.Send(&OutPacket{
		Type: "error",
		Data: map[string]string{
			"message": msg,
		},
	})
}

// Pinger sends ping requests to the client periodically
func (c *Client) Pinger() {
	for {
		c.Send(&OutPacket{
			Type: "ping",
			Data: nil,
		})
		time.Sleep(20 * time.Second)
		if time.Since(c.LastPong) > 20*time.Second {
			// TODO : add mission-critical guild de-registration here
			logrus.Debug("closing socket : ping timeout")
			if err := c.Conn.Close(); err != nil {
				logrus.Warn("error closing socket", err)
			}
		}
	}
}

// Reader eternally waits for things to read from the client
func (c *Client) Reader() {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			logrus.Warn("Error reading message from client", err)
			continue
		}
		var p Packet
		if err := json.Unmarshal(msg, &p); err != nil {
			logrus.Warn("Error parsing client packet", err)
			continue
		}
		if c.Bus[p.Type] != nil {
			c.Bus[p.Type].Handler(*c, c.Bus[p.Type], p.Data)
		}
	}
}

// Writer eternally waits for things to write to the client
func (c *Client) Writer() {
	for {
		msg := <-c.Out // wait for a new message to be sent
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			logrus.Warnf("Error writing data to events : %v", err)
			logrus.Debugf("Closing Socket : Data write error")
			_ = c.Conn.Close()
			// TODO : add mission-critical guild de-registration here
			return
		}
	}
}
