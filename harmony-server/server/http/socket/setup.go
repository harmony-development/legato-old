package socket

import (
	"harmony-server/server/http/socket/events"
	"harmony-server/server/state/event"
	"time"
)

func (h Handler) Setup() {
	e := events.NewEvents(h.DB, h.State)
	h.Bus.Bind(event.NewEvent(e.Ping, "ping", 3 * time.Second, 10))
	h.Bus.Bind(event.NewEvent(e.Subscribe, "subscribe", 3 * time.Second, 10))
}