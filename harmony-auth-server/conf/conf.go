package conf

import (
	"fmt"
	"regexp"
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
)

var (
	// UsernameLenMessage is the message printed when the username is out of bounds
	UsernameLenMessage = fmt.Sprintf("username must be between %v and %v characters long", UsernameLenMin, UsernameLenMax)

	// PasswordLenMessage is the message printed when the password is out of bounds
	PasswordLenMessage = fmt.Sprintf("password must be between %v and %v characters long", PassLenMin, PassLenMax)

	// EmailValidation is a regex pattern to verify if an email is valid
	EmailValidation = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)