package v1

import (
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/db/queries"
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

type GetGuildListResponse struct {
	Guilds []queries.GetGuildListRow `json:"guilds"`
}
