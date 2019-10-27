package socket

import (
	"encoding/json"
)

// EventHandler is a type of function that handles an event
type EventHandler func(data interface{}, ws *WebSocket)

// Event is the structure the client sends to the server for handling
type Event struct {
	Name string      // The event we want to listen to
	Data interface{} // The data that is received at that event
}

// ParseMessage parses a raw message to the server into an Event
func ParseMessage(unparsed []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(unparsed, event)
	return event, err
}

// Raw converts things back to byte to send to clients
func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
