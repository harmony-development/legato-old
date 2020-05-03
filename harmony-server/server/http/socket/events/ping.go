package events

import (
	"encoding/json"
	"harmony-server/server/http/socket/handling"
	"time"
)

// Ping handles ping requests from the client
func (e Events) Ping(ws handling.Client, event *handling.Event, _ *json.RawMessage) {
	if !event.Limiter.Allow() {
		return
	}
	ws.LastPong = time.Now()
}