// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package authv1impl

import dynamicauth "github.com/harmony-development/legato/dynamic_auth"

var initialStep = dynamicauth.NewChoiceStep(
	[]string{
		"login",
		"register",
		"other-options",
	}, "initial-step", false,
)

// nolint
// this is a protocol-level constant
var loginStep = dynamicauth.NewFormStep(
	[]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
		{Name: "password", FieldType: "password"},
	}, "login", true,
)

// nolint
// this is a protocol-level constant
var registerStep = dynamicauth.NewFormStep(
	[]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
		{Name: "username", FieldType: "username"},
		{Name: "password", FieldType: "new-password"},
	}, "register", true,
)

// nolint
// this is a protocol-level constant
var otherOptionsStep = dynamicauth.NewChoiceStep(
	[]string{
		"reset-password",
	}, "other-options", true,
)

// nolint
// this is a protocol-level constant
var resetPasswordStep = dynamicauth.NewFormStep(
	[]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
	}, "reset-password", true,
)
