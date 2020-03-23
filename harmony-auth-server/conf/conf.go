package conf

import (
	"fmt"
	"regexp"
	"time"
)

const (
	// InstanceAPIVersion specifies the target version of instance API to request from (/api/vX/.....)
	InstanceAPIVersion = "v1"
	// UsernameLenMin is the min length for a username (1 letter usernames would be kinda insane)
	UsernameLenMin = 2
	// UsernameLenMax is the max length for a username (if it were longer the UI would look ridiculous)
	UsernameLenMax = 48
	// PassLenMin is the min length for a password (let the frontend do any other heavy checks)
	PassLenMin = 5
	// PassLenMax is the max length for a password (maybe it should be 256?)
	PassLenMax = 128
	// SessionLength is the char length of characters in the session generated
	SessionLength = 16
	// SessionExpire says how far into the future the session should still be valid
	SessionExpire = 24 * time.Hour
	// SessionCacheMax is the max capacity of the session LRU cache
	SessionCacheMax = 500000
	// IDCacheMax is the max capacity of the userid LRU cache
	IDCacheMax = 500000
	// AvatarQuality is the quality to downscale to for avatars
	AvatarQuality = 60
	// AvatarWidth is the width of avatars for downscaling
	AvatarWidth = 128
	// AvatarHeight is the height of avatars for downscaling
	AvatarHeight = 128
	// AvatarCrop determines whether to crop when resizing an image. True is recommended. Otherwise, it looks silly
	AvatarCrop = true
)

var (
	// UsernameLenMessage is the message printed when the username is out of bounds
	UsernameLenMessage = fmt.Sprintf("username must be between %v and %v characters long", UsernameLenMin, UsernameLenMax)

	// PasswordLenMessage is the message printed when the password is out of bounds
	PasswordLenMessage = fmt.Sprintf("password must be between %v and %v characters long", PassLenMin, PassLenMax)

	// EmailValidation is a regex pattern to verify if an email is valid
	EmailValidation = regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	)
)
