package events

import (
	"encoding/json"
	"harmony-server/server/state/event"
	"harmony-server/server/state/guild"
	"time"
)

// Ping handles ping requests from the client
func (e Events) Ping(ws guild.Client, event *event.Event, _ *json.RawMessage) {
	if !event.Limiter.Allow() {
		return
	}
	ws.LastPong = time.Now()
}