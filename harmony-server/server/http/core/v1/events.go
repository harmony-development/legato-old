package v1

import (
	"encoding/json"
)

const (
	ActionEventType        = "action"
	ChannelCreateEventType = "channel_create"
	ChannelDeleteEventType = "channel_delete"
	AvatarUpdateEventType  = "avatar_update"
	GuildDeleteEventType   = "guild_delete"
	GuildUpdateEventType   = "guild_update"
	MessageDeleteEventType = "message_delete"
	MessageCreateEventType = "message_create"
	MessageUpdateEventType = "message_update"
	UserUpdateEventType    = "user_update"
)

type MessageUpdateFlags uint64

const (
	UpdateContent MessageUpdateFlags = 1 << iota
	UpdateActions
	UpdateEmbeds
)

// ActionEvent is the data that will be sent to a client on an action trigger
type ActionEvent struct {
	GuildID   uint64 `json:"guild_id"`
	ChannelID uint64 `json:"channel_id"`
	MessageID uint64 `json:"message_id"`
	TriggerID uint64 `json:"trigger_id"`
	Action    struct {
		ID   string `json:"id"`
		Data string `json:"data"`
	} `json:"action"`
}

// ChannelCreateEvent is the data that will be sent to a client on a channel create
type ChannelCreateEvent struct {
	GuildID     uint64 `json:"guild_id"`
	ChannelName string `json:"channel_name"`
	ChannelID   uint64 `json:"channel_id"`
}

// ChannelDeleteEvent is the data that will be sent to a client on a channel delete
type ChannelDeleteEvent struct {
	ChannelID uint64 `json:"channel_id"`
	GuildID   uint64 `json:"guild_id"`
}

// AvatarUpdateEvent is the data that will be sent to a client on an avatar update
type AvatarUpdateEvent struct {
	UserID    uint64 `json:"user_id"`
	NewAvatar string `json:"new_avatar"`
}

// GuildDeleteEvent is the data that will be sent to a client on a guild delete
type GuildDeleteEvent struct {
	GuildID uint64 `json:"guild_id"`
}

// MessageDeleteEvent is the data that will be sent to a client on a message delete
type MessageDeleteEvent struct {
	GuildID   uint64 `json:"guild_id"`
	ChannelID uint64 `json:"channel_id"`
	MessageID uint64 `json:"message_id"`
}

// MessageCreateEvent is the data that will be sent to a client on a message create
type MessageCreateEvent struct {
	GuildID     uint64            `json:"guild_id"`
	ChannelID   uint64            `json:"channel_id"`
	CreatedAt   int64             `json:"created_at"`
	Message     string            `json:"message"`
	Attachments []string          `json:"attachments,omitempty"`
	AuthorID    uint64            `json:"author_id"`
	MessageID   uint64            `json:"message_id"`
	Actions     []json.RawMessage `json:"actions,omitempty"`
	Embeds      []json.RawMessage `json:"embeds,omitempty"`
}

// GuildUpdateEvent is the data that will be sent to a client on a guild update
type GuildUpdateEvent struct {
	GuildID uint64 `json:"guild_id"`
	Name    string `json:"name,omitempty"`
	Picture string `json:"picture,omitempty"`
}

// UserUpdateEvent is the data that will be sent to a client on a user update
type UserUpdateEvent struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username,omitempty"`
}

// MessageUpdateEvent is the data that will be sent to a client on a message update
type MessageUpdateEvent struct {
	GuildID   uint64             `json:"guild_id"`
	ChannelID uint64             `json:"channel_id"`
	MessageID uint64             `json:"message_id"`
	Flags     MessageUpdateFlags `json:"flags"`
	EditedAt  int64              `json:"edited_at"`
	Message   string             `json:"message,omitempty"`
	Actions   []json.RawMessage  `json:"actions,omitempty"`
	Embeds    []json.RawMessage  `json:"embeds,omitempty"`
}
