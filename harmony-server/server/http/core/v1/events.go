package v1

const (
	ActionEventType        = "action"
	ChannelCreateEventType = "channelCreate"
	ChannelDeleteEventType = "channelDelete"
	AvatarUpdateEventType  = "avatarUpdate"
	GuildDeleteEventType   = "guildDelete"
	MessageDeleteEventType = "messageDelete"
)

// ActionEvent is the data that will be sent to a client on an action trigger
type ActionEvent struct {
	GuildID   uint64 `json:"guildID"`
	ChannelID uint64 `json:"channelID"`
	MessageID uint64 `json:"messageID"`
	TriggerID uint64 `json:"triggerID"`
	Action    struct {
		ID   string `json:"id"`
		Data string `json:"data"`
	} `json:"action"`
}

// ChannelCreateEvent is the data that will be sent to a client on a channel create
type ChannelCreateEvent struct {
	GuildID     uint64 `json:"guildID"`
	ChannelName string `json:"channelName"`
	ChannelID   uint64 `json:"channelID"`
}

// ChannelDeleteEvent is the data that will be sent to a client on a channel delete
type ChannelDeleteEvent struct {
	ChannelID uint64 `json:"channelID"`
	GuildID   uint64 `json:"guildID"`
}

// AvatarUpdateEvent is the data that will be sent to a client on an avatar update
type AvatarUpdateEvent struct {
	UserID    uint64 `json:"userID"`
	NewAvatar string `json:"newAvatar"`
}

// GuildDeleteEvent is the data that will be sent to a client on a guild delete
type GuildDeleteEvent struct {
	GuildID uint64 `json:"guildID"`
}

// MessageDeleteEvent is the data that will be sent to a client on a message delete
type MessageDeleteEvent struct {
	GuildID   uint64 `json:"guildID"`
	ChannelID uint64 `json:"channelID"`
	MessageID uint64 `json:"messageID"`
}
