package v1

import (
	"harmony-server/server/db"
)

type UserInfoResponse struct {
	UserName   string        `json:"user_name"`
	UserAvatar string        `json:"user_avatar"`
	UserStatus db.UserStatus `json:"user_status"`
	GuildList  string        `json:"guild_list"`
}
