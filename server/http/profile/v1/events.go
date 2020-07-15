package v1

import (
	profilev1 "github.com/harmony-development/legato/gen/profile"
)

const (
	AvatarUpdateEventType = "avatar_update"
	UserUpdateEventType   = "user_update"
)

// AvatarUpdateEvent is the data that will be sent to a client on an avatar update
type AvatarUpdateEvent struct {
	UserID    string `json:"user_id"`
	NewAvatar string `json:"new_avatar"`
}

// UsernameUpdateEvent is the data that will be sent to a client on a username update
type UsernameUpdateEvent struct {
	UserID   string `json:"user_id"`
	Username string `json:"username,omitempty"`
}

// StatusUpdateEvent is the data that will be sent to a client on a user status update
type StatusUpdateEvent struct {
	UserID string               `json:"user_id"`
	Status profilev1.UserStatus `json:"status"`
}
