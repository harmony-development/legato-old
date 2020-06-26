package v1

type GuildCreateResponse struct {
	GuildID string `json:"guild_id"`
}

type GuildInfoResponse struct {
	GuildName    string `json:"guild_name"`
	GuildOwner   string `json:"guild_owner"`
	GuildPicture string `json:"guild_picture"`
}

type MemberListResponse struct {
	Members []string `json:"members"`
}

type JoinedGuildsResponseGuild struct {
	GuildID      string `json:"guild_id"`
	GuildName    string `json:"guild_name"`
	GuildOwner   string `json:"guild_owner"`
	GuildPicture string `json:"guild_picture"`
}

type JoinedGuildsResponse struct {
	Guilds []JoinedGuildsResponseGuild `json:"guilds"`
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
	Actions   []Action `json:"actions,omitempty"`
}

type ChannelCreateResponse struct {
	GuildID     string `json:"guild_id"`
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
}

type MessageCreateResponse struct {
	MessageID string `json:"message_id"`
}

type MessageListResponse struct {
	Messages []Message `json:"messages"`
}

type LoginResponse struct {
	UserID  string `json:"user_id"`
	Session string `json:"session"`
}

type InviteCreateResponse struct {
	Name string `json:"invite_name"`
	ID   string `json:"invite_id"`
	Uses int32  `json:"invite_uses"`
}

type Invite struct {
	ID        string `json:"invite_id"`
	GuildID   string `json:"guild_id"`
	Uses      int32  `json:"invite_uses,omitempty"`
	UsedCount int32  `json:"invite_used"`
}

type GetInvitesResponse struct {
	Invites []Invite `json:"invites"`
}

type JoinGuildResponse struct {
	GuildID string `json:"guild_id"`
}

type RegisterResponse LoginResponse
