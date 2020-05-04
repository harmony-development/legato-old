package socket

import (
	"harmony-server/server/http/socket/client"
	"harmony-server/server/http/socket/events"
	"time"
)

func (h Handler) Setup() {
	e := events.NewEvents(h.DB, h.State)
	h.Bus.Bind(client.NewEvent(e.Ping, "ping", 3 * time.Second, 10))
	h.Bus.Bind(client.NewEvent(e.Subscribe, "subscribe", 3 * time.Second, 10))
}