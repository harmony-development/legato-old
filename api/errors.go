package api

import harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"

const (
	InternalServerError = "h.internal-server-error"

	Other = "h.other"
)

type Error harmonytypesv1.Error

func (e *Error) Error() string {
	return e.HumanMessage
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
