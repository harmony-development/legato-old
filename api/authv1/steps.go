package authv1impl

import dynamicauth "github.com/harmony-development/legato/dynamic_auth"

var initialStep = dynamicauth.NewChoiceStep(
	[]string{
		"login",
		"register",
		"other-options",
	}, "initial-step", false,
)

var loginStep = dynamicauth.NewFormStep(
	[]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
		{Name: "password", FieldType: "password"},
	}, "login", true,
)

var registerStep = dynamicauth.NewFormStep(
	[]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
		{Name: "username", FieldType: "username"},
		{Name: "password", FieldType: "new-password"},
	}, "register", true,
)

var otherOptionsStep = dynamicauth.NewChoiceStep(
	[]string{
		"reset-password",
	}, "other-options", true,
)

var resetPasswordStep = dynamicauth.NewFormStep(
	[]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
	}, "reset-password", true,
)
