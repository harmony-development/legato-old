package v1

const (
	AvatarUpdateEventType = "avatar_update"
	UserUpdateEventType   = "user_update"
)

// AvatarUpdateEvent is the data that will be sent to a client on an avatar update
type AvatarUpdateEvent struct {
	UserID    string `json:"user_id"`
	NewAvatar string `json:"new_avatar"`
}

// UserUpdateEvent is the data that will be sent to a client on a user update
type UserUpdateEvent struct {
	UserID   string `json:"user_id"`
	Username string `json:"username,omitempty"`
}
