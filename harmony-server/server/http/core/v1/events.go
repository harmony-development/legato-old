package v1

const (
	ActionEventType        = "action"
	ChannelCreateEventType = "channel_create"
	ChannelDeleteEventType = "channel_delete"
	AvatarUpdateEventType  = "avatar_update"
	GuildDeleteEventType   = "guild_delete"
	MessageDeleteEventType = "message_delete"
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
