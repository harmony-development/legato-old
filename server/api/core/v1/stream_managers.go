package v1

import corev1 "github.com/harmony-development/legato/gen/core"

type ActionSubscriptionManager interface {
	Subscribe(userID uint64, stream corev1.CoreService_StreamEventsServer) chan struct{}
	Unsubscribe(userID uint64, stream corev1.CoreService_StreamEventsServer)
	Broadcast(to uint64, action *corev1.Event)
}

type GuildSubscriptionManager interface {
	Subscribe(guildID, userID uint64, server corev1.CoreService_StreamEventsServer) chan struct{}
	UnsubscribeUser(userID uint64)
	UnsubscribeGuild(guildID uint64)
	UnsubscribeUserFromGuild(userID, guildID uint64)
	Broadcast(to uint64, event *corev1.Event)
}

type HomeserverSubscriptionManager interface {
	Subscribe(userID uint64, s corev1.CoreService_StreamEventsServer) chan struct{}
	Unsubscribe(userID uint64, s corev1.CoreService_StreamEventsServer)
	Broadcast(userID uint64, e *corev1.Event)
}

type SubscriptionManager struct {
	Actions    ActionSubscriptionManager
	Guild      GuildSubscriptionManager
	Homeserver HomeserverSubscriptionManager
}
