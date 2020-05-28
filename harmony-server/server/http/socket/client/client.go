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
	PingTicker *time.Ticker
	Out        chan []byte
	Deregister func(client *Client)
}

// Packet is the standard way all messages are delivered and received from the server
type Packet struct {
	Type string           `json:"type"`
	Data *json.RawMessage `json:"data"`
}

// OutPacket is like packet but the data doesn't necessarily have to be raw bytes
type OutPacket struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
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

// Reader eternally waits for things to read from the client
func (c *Client) Reader() {
	defer func() {
		logrus.Debug("Reader routine exited")
	}()
	c.Conn.SetReadLimit(4096)
	if err := c.Conn.SetReadDeadline(time.Now().Add(1 * time.Minute)); err != nil {
		logrus.Warn(err)
	}
	c.Conn.SetPongHandler(func(string) error {
		if err := c.Conn.SetReadDeadline(time.Now().Add(1 * time.Minute)); err != nil {
			logrus.Warn(err)
		}
		return nil
	})
	for {
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			logrus.Warn("Read failure: ", err)
			if c.UserID != nil {
				c.Deregister(c)
			}
			c.Conn.Close()
			return
		}
		if msgType == websocket.TextMessage {
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
}

// Writer eternally waits for things to write to the client
func (c *Client) Writer() {
	c.PingTicker = time.NewTicker(15 * time.Second)
	defer func() {
		c.PingTicker.Stop()
		logrus.Debug("Writer routine exited")
	}()
	for {
		select {
		case msg := <-c.Out:
			err := c.Conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				logrus.Warnf("Error writing data to events : %v", err)
				logrus.Debugf("Closing Socket : Data write error")
				if c.UserID != nil {
					c.Deregister(c)
				}
				_ = c.Conn.Close()
				return
			}
		case <-c.PingTicker.C:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(15 * time.Second)); err != nil {
				logrus.Warn(err)
			}
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logrus.Debug("Write failure: ", err)
				if c.UserID != nil {
					c.Deregister(c)
				}
				c.Conn.Close()
				return
			}
		}
	}
}
