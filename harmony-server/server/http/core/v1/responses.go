package v1

type GuildCreateResponse struct {
	GuildID string `json:"guild_id"`
}

type GuildInfoResponse struct {
	GuildName    string `json:"guild_name"`
	GuildOwner   string `json:"guild_owner"`
	GuildPicture string `json:"guild_picture"`
}

type UserInfoResponse struct {
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
}

type MemberListResponse struct {
	Members []string `json:"members"`
}

type ChannelListResponse struct {
	Channels []Channel `json:"channels"`
}

type Message struct {
	MessageID string   `json:"message_id"`
	GuildID   string   `json:"guild_id"`
	ChannelID string   `json:"channel_id"`
	AuthorID  string   `json:"author_id"`
	CreatedAt string   `json:"created_at"`
	EditedAt  *string  `json:"edited_at,omitempty"`
	Content   string   `json:"content"`
	Embeds    []Embed  `json:"embeds,omitempty"`
	Actions   []Action `json:"actions",omitempty`
}

type MessageListResponse struct {
	Messages []Message `json:"messages"`
}

type LoginResponse struct {
	Session string
}

type RegisterResponse LoginResponse
