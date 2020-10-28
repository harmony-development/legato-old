package v1

import corev1 "github.com/harmony-development/legato/gen/core"

type ActionSubscriptionManager interface {
	Subscribe(userID uint64, stream corev1.CoreService_StreamActionEventsServer) chan struct{}
	Unsubscribe(userID uint64, stream corev1.CoreService_StreamActionEventsServer)
	Broadcast(to uint64, action *corev1.ActionEvent)
}

type GuildSubscriptionManager interface {
	Subscribe(guildID, userID uint64, server corev1.CoreService_StreamGuildEventsServer) chan struct{}
	UnsubscribeUser(userID uint64)
	UnsubscribeGuild(guildID uint64)
	UnsubscribeUserFromGuild(userID, guildID uint64)
	Broadcast(to uint64, event *corev1.GuildEvent)
}

type HomeserverSubscriptionManager interface {
	Subscribe(userID uint64, s corev1.CoreService_StreamHomeserverEventsServer) chan struct{}
	Unsubscribe(userID uint64, s corev1.CoreService_StreamHomeserverEventsServer)
	Broadcast(userID uint64, e *corev1.HomeserverEvent)
}

type SubscriptionManager struct {
	Actions    ActionSubscriptionManager
	Guild      GuildSubscriptionManager
	Homeserver HomeserverSubscriptionManager
}
