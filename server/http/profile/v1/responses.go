package v1

import (
	"github.com/harmony-development/legato/server/db"
)

type UserInfoResponse struct {
	UserName   string        `json:"user_name"`
	UserAvatar string        `json:"user_avatar"`
	UserStatus db.UserStatus `json:"user_status"`
	GuildList  string        `json:"guild_list,omitempty"`
}

type GetUserMetadataResponse struct {
	Metadata string `json:"metadata"`
}

type GetGuildListGuild struct {
	GuildID    string `json:"guild_id"`
	HomeServer string `json:"home_server"`
}

type GetGuildListResponse struct {
	Guilds []GetGuildListGuild `json:"guilds"`
}
