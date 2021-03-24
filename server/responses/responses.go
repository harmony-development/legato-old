package responses

import harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"

// basically a list of API responses for i18n compatibility

const (
	TooManyRequests     = "h.too-many-requests"
	InternalServerError = "h.internal-server-error"
	BadRequest          = "h.bad-request"

	NoSessionProvided = "h.blank-session"
	BadSession        = "h.bad-session"

	MissingGuildID   = "h.missing-guild-id"
	MissingChannelID = "h.missing-channel-id"
	MissingMessageID = "h.missing-message-id"
	MissingUserID    = "h.bad-user-id"
	MissingAuthID    = "h.missing-auth-id"
	MissingFilename  = "h.missing-file-name"
	MissingFiles     = "h.missing-files"

	BadGuildID   = "h.bad-guild-id"
	BadChannelID = "h.bad-channel-id"
	BadMessageID = "h.bad-message-id"
	BadUserID    = "h.bad-user-id"
	BadAuthID    = "h.bad-auth-id"

	NotAuthor = "h.not-author"

	NotJoined = "h.not-joined"

	IsOwner              = "h.is-owner"
	NotOwner             = "h.not-owner"
	NotEnoughPermissions = "h.not-enough-permissions"

	AlreadyRegistered = "h.already-registered"
	AlreadyInGuild    = "h.already-in-guild"

	BadEmail    = "h.bad-email"
	BadPassword = "h.bad-password"
	BadUsername = "h.bad-username"
	NameTooLong = "h.name-too-long"

	IncorrectPassword = "h.incorrect-password"

	EntirelyBlank = "h.entirely-blank"

	BadAction = "h.bad-action"

	NoMetadata = "h.no-metadata"

	Other = "h.other"

	BadChoice = "h.bad-auth-choice"

	MissingForm = "h.missing-form"

	TooManyFiles = "h.too-many-files"

	BannedFromGuild = "h.banned-from-guild"
)

// Error is a wrapper around harmonytypesv1.Error implementing the Error interface
type Error harmonytypesv1.Error

func (e *Error) Error() string {
	return e.Identifier
}

func NewError(code string) error {
	return &Error{
		Identifier: code,
	}
}

func NewOther(msg string) error {
	return &Error{
		Identifier:   Other,
		HumanMessage: msg,
	}
}
