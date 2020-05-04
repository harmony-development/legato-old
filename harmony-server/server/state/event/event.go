package event

import (
	"encoding/json"
	"golang.org/x/time/rate"
	"harmony-server/server/state/guild"
	"time"
)

// Handler is a function that is called for an event
type Handler func(c guild.Client, e *Event, data *json.RawMessage)

// Bus is a collection of websocket events
type Bus map[string]*Event

// Event is a handler for websocket messages
type Event struct {
	Handler Handler
	Path    string
	Limiter *rate.Limiter
}

// NewEvent creates a new socket event handler
func NewEvent(handler Handler, path string, rateLimit time.Duration, burst int) Event {
	return Event{
		Handler: handler,
		Path:    path,
		Limiter: rate.NewLimiter(rate.Every(rateLimit), burst),
	}
}

// Bind is syntactic sugar for simply setting a value in the map
func (b Bus) Bind(e Event) {
	b[e.Path] = &e
}
