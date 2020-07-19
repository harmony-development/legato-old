package v1

import corev1 "github.com/harmony-development/legato/gen/core"

type GuildID uint64
type UserID uint64

type StreamState struct {
	GuildState
	ActionState
}

var streamState = StreamState{
	GuildState: GuildState{
		guildEvents: map[UserID]map[GuildID][]corev1.CoreService_StreamGuildEventsServer{},
		subs:        map[GuildID]map[UserID]struct{}{},
	},
	ActionState: ActionState{
		actionEvents: map[UserID][]corev1.CoreService_StreamActionEventsServer{},
	},
}
