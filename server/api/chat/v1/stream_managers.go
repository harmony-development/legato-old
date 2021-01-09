package v1

import (
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
)

type StreamManager interface {
	RegisterClient(userID uint64, s chatv1.ChatService_StreamEventsServer) chan struct{}

	AddGuildSubscription(s chatv1.ChatService_StreamEventsServer, to uint64)
	AddHomeserverSubscription(s chatv1.ChatService_StreamEventsServer)
	AddActionSubscription(s chatv1.ChatService_StreamEventsServer)

	RemoveGuildSubscription(userID, guildID uint64)

	BroadcastGuild(to uint64, event *chatv1.Event)
	BroadcastHomeserver(userid uint64, event *chatv1.Event)
	BroadcastAction(userid uint64, event *chatv1.Event)
}
