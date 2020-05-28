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

type LoginResponse struct {
	Session string
}

type RegisterResponse LoginResponse
