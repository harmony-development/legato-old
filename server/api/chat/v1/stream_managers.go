package v1

import chatv1 "github.com/harmony-development/legato/gen/chat/v1"

type ActionSubscriptionManager interface {
	Subscribe(userID uint64, stream chatv1.ChatService_StreamEventsServer) chan struct{}
	Unsubscribe(userID uint64, stream chatv1.ChatService_StreamEventsServer)
	Broadcast(to uint64, action *chatv1.Event)
}

type GuildSubscriptionManager interface {
	Subscribe(guildID, userID uint64, server chatv1.ChatService_StreamEventsServer) chan struct{}
	UnsubscribeUser(userID uint64)
	UnsubscribeGuild(guildID uint64)
	UnsubscribeUserFromGuild(userID, guildID uint64)
	Broadcast(to uint64, event *chatv1.Event)
}

type HomeserverSubscriptionManager interface {
	Subscribe(userID uint64, s chatv1.ChatService_StreamEventsServer) chan struct{}
	Unsubscribe(userID uint64, s chatv1.ChatService_StreamEventsServer)
	Broadcast(userID uint64, e *chatv1.Event)
}

type SubscriptionManager struct {
	Actions    ActionSubscriptionManager
	Guild      GuildSubscriptionManager
	Homeserver HomeserverSubscriptionManager
}
