package events

import (
	"encoding/json"

	"github.com/harmony-development/legato/server/http/socket/client"
)

func (e Events) SubscribeToUserUpdates(ws client.Client, _ *client.Event, _ *json.RawMessage) {
	e.State.UserUpdateListeners[&ws] = struct{}{}
}
