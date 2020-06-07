package v1

import "harmony-server/server/db/queries"

type UserInfoResponse struct {
	UserName   string             `json:"user_name"`
	UserAvatar string             `json:"user_avatar"`
	UserStatus queries.Userstatus `json:"user_status"`
}
