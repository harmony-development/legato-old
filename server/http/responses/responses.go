package responses

// basically a list of API responses for i18n compatibility

const (
	InvalidEmail           = "auth.invalid-email"
	AlreadyRegistered      = "auth.already-registered"
	InvalidPassword        = "auth.invalid-password"
	InvalidSession         = "invalid-session"
	TooManyRequests        = "too-many-requests"
	InvalidRequest         = "invalid-request"
	InsufficientPrivileges = "insufficient-privileges"
	UserNotFound           = "user-not-found"
	GuildNotFound          = "guild.not-found"
	NotInGuild             = "guild.not-member"
	MetadataNotFound       = "user.metadata-not-found"
	NonceNotFound          = "nonce-not-found"
	MissingLocationGuild   = "missing-location-guild"
	MissingLocationChannel = "missing-location-channel"
	MissingLocationMessage = "missing-location-message"
	BadLocationGuild       = "invalid-location-guild"
	BadLocationChannel     = "invalid-location-channel"
	BadLocationMessage     = "invalid-location-message"
	InternalServerError    = "internal-server-error"
	TeaPot                 = "i-am-a-teapot-and-will-not-serve-coffee"
	UnknownError           = "unknown"
)

type WithFields struct {
	Message string      `json:"message"`
	Fields  interface{} `json:"fields"`
}

func UsernameLength(minLength, maxLength int) WithFields {
	return WithFields{
		Message: "register.username-length",
		Fields: map[string]interface{}{
			"minLength": minLength,
			"maxLength": maxLength,
		},
	}
}

func PasswordLength(minLength, maxLength int) WithFields {
	return WithFields{
		Message: "register.password-length",
		Fields: map[string]interface{}{
			"minLength": minLength,
			"maxLength": maxLength,
		},
	}
}

func PasswordPolicy(minUpper, minLower, minNumbers, minSpecial int) WithFields {
	return WithFields{
		Message: "register.password-policy",
		Fields: map[string]interface{}{
			"minUpper":   minUpper,
			"minLower":   minLower,
			"minNumbers": minNumbers,
			"minSpecial": minSpecial,
		},
	}
}
