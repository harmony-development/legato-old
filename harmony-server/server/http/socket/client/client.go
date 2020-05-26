package client

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Handler is a function that is called for an event
type Handler func(c Client, e *Event, data *json.RawMessage)

// Client is the data structure for a connected client
type Client struct {
	*sync.RWMutex
	Conn       *websocket.Conn
	Bus        Bus
	UserID     *uint64
	LastPong   time.Time
	Out        chan []byte
	Deregister func(client *Client)
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
		if c.Conn == nil {
			return
		}
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
			c.Deregister(c)
			return
		}
	}
}
