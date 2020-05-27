package events

import (
	"encoding/json"

	"harmony-server/server/http/socket/client"
)

func (e Events) SubscribeToUserUpdates(ws client.Client, _ *client.Event, _ *json.RawMessage) {
	e.State.UserUpdateListeners[&ws] = struct{}{}
}