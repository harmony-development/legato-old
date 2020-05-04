package events

import (
	"encoding/json"
	"harmony-server/server/http/socket/client"
	"time"
)

// Ping handles ping requests from the client
func (e Events) Ping(ws client.Client, event *client.Event, _ *json.RawMessage) {
	if !event.Limiter.Allow() {
		return
	}
	ws.LastPong = time.Now()
}
