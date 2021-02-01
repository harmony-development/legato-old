package v1

import (
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
)

type StreamManager interface {
	RegisterClient(userID uint64, s chan *chatv1.Event, done chan struct{})

	AddGuildSubscription(s chan *chatv1.Event, to uint64)
	AddHomeserverSubscription(s chan *chatv1.Event)
	AddActionSubscription(s chan *chatv1.Event)

	RemoveGuildSubscription(userID, guildID uint64)

	BroadcastGuild(to uint64, event *chatv1.Event)
	BroadcastHomeserver(userid uint64, event *chatv1.Event)
	BroadcastAction(userid uint64, event *chatv1.Event)
}
