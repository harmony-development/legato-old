package v1

import (
	profilev1 "github.com/harmony-development/legato/gen/profile"
)

type UserInfoResponse struct {
	UserName   string               `json:"user_name"`
	UserAvatar string               `json:"user_avatar"`
	UserStatus profilev1.UserStatus `json:"user_status"`
	GuildList  string               `json:"guild_list,omitempty"`
}

type GetUserMetadataResponse struct {
	Metadata string `json:"metadata"`
}
