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
	GuildNotFound          = "guild.not-found"
	UnknownError           = "unknown"
)

type WithFields struct {
	message string
	fields  interface{}
}

func UsernameLength(minLength, maxLength int) WithFields {
	return WithFields{
		message: "register.username-length",
		fields: map[string]interface{}{
			"minLength": minLength,
			"maxLength": maxLength,
		},
	}
}

func PasswordLength(minLength, maxLength int) WithFields {
	return WithFields{
		message: "register.password-length",
		fields: map[string]interface{}{
			"minLength": minLength,
			"maxLength": maxLength,
		},
	}
}

func PasswordPolicy(minUpper, minLower, minNumbers, minSpecial int) WithFields {
	return WithFields{
		message: "register.password-policy",
		fields: map[string]interface{}{
			"minUpper":   minUpper,
			"minLower":   minLower,
			"minNumbers": minNumbers,
			"minSpecial": minSpecial,
		},
	}
}
