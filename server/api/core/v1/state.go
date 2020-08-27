package v1

import corev1 "github.com/harmony-development/legato/gen/core"

type (
	GuildID uint64
	UserID  uint64
)

type StreamState struct {
	GuildState
	ActionState
}

var streamState = StreamState{
	GuildState: GuildState{
		guildEvents:    make(map[UserID]map[GuildID][]corev1.CoreService_StreamGuildEventsServer),
		serverChannels: make(map[corev1.CoreService_StreamGuildEventsServer]chan struct{}),
		subs:           make(map[GuildID]map[UserID]struct{}),
	},
	ActionState: ActionState{
		actionEvents: make(map[UserID][]corev1.CoreService_StreamActionEventsServer),
	},
}
