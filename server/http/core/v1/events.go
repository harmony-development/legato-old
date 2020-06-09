package v1

import (
	"encoding/json"
)

const (
	ActionEventType        = "action"
	ChannelCreateEventType = "channel_create"
	ChannelDeleteEventType = "channel_delete"
	GuildDeleteEventType   = "guild_delete"
	GuildUpdateEventType   = "guild_update"
	MessageDeleteEventType = "message_delete"
	MessageCreateEventType = "message_create"
	MessageUpdateEventType = "message_update"
)

type MessageUpdateFlags uint64

const (
	UpdateContent MessageUpdateFlags = 1 << iota
	UpdateActions
	UpdateEmbeds
)

// ActionEvent is the data that will be sent to a client on an action trigger
type ActionEvent struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
	TriggerID string `json:"trigger_id"`
	Action    struct {
		ID   string `json:"id"`
		Data string `json:"data"`
	} `json:"action"`
}

// ChannelCreateEvent is the data that will be sent to a client on a channel create
type ChannelCreateEvent struct {
	GuildID     string `json:"guild_id"`
	ChannelName string `json:"channel_name"`
	ChannelID   string `json:"channel_id"`
}

// ChannelDeleteEvent is the data that will be sent to a client on a channel delete
type ChannelDeleteEvent struct {
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
}

// GuildDeleteEvent is the data that will be sent to a client on a guild delete
type GuildDeleteEvent struct {
	GuildID string `json:"guild_id"`
}

// MessageDeleteEvent is the data that will be sent to a client on a message delete
type MessageDeleteEvent struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
}

// MessageCreateEvent is the data that will be sent to a client on a message create
type MessageCreateEvent struct {
	GuildID     string            `json:"guild_id"`
	ChannelID   string            `json:"channel_id"`
	CreatedAt   int64             `json:"created_at"`
	Message     string            `json:"message"`
	Attachments []string          `json:"attachments,omitempty"`
	AuthorID    string            `json:"author_id"`
	MessageID   string            `json:"message_id"`
	Actions     []json.RawMessage `json:"actions,omitempty"`
	Embeds      []json.RawMessage `json:"embeds,omitempty"`
}

// GuildUpdateEvent is the data that will be sent to a client on a guild update
type GuildUpdateEvent struct {
	GuildID string `json:"guild_id"`
	Name    string `json:"name,omitempty"`
	Picture string `json:"picture,omitempty"`
}

// MessageUpdateEvent is the data that will be sent to a client on a message update
type MessageUpdateEvent struct {
	GuildID   string             `json:"guild_id"`
	ChannelID string             `json:"channel_id"`
	MessageID string             `json:"message_id"`
	Flags     MessageUpdateFlags `json:"flags"`
	EditedAt  int64              `json:"edited_at"`
	Message   string             `json:"message,omitempty"`
	Actions   []json.RawMessage  `json:"actions,omitempty"`
	Embeds    []json.RawMessage  `json:"embeds,omitempty"`
}
