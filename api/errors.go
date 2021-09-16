// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package api

import harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"

const (
	ErrorBadAuthID      = "h.bad-auth-id"
	ErrorBadChoice      = "h.bad-auth-choice"
	ErrorBadFormData    = "h.bad-form-data"
	ErrorBadCredentials = "h.bad-credentials"
	ErrorBadStep        = "h.bad-step"

	ErrorInternalServerError = "h.internal-server-error"
	ErrorOther               = "h.other"
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
		Identifier:   ErrorOther,
		HumanMessage: msg,
	}
}
