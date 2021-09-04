package authv1impl

import (
	"context"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	dynamicauth "github.com/harmony-development/legato/server/dynamic_auth"
	"github.com/harmony-development/legato/server/key"
)

type AuthV1 struct {
	authv1.DefaultAuthService
	keyManager key.KeyManager
}

var steps = []dynamicauth.Step{
	dynamicauth.NewChoiceStep([]string{
		"login",
		"register",
		"other-options",
	}, "initial-step", false),

	dynamicauth.NewFormStep([]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
		{Name: "password", FieldType: "password"},
	}, "login", true),
	dynamicauth.NewFormStep([]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
		{Name: "username", FieldType: "username"},
		{Name: "password", FieldType: "new-password"},
	}, "register", true),

	dynamicauth.NewChoiceStep([]string{
		"reset-password",
	}, "other-options", true),
	dynamicauth.NewFormStep([]dynamicauth.FormField{
		{Name: "email", FieldType: "email"},
	}, "reset-password", true),
}

var rawSteps = map[string]*authv1.AuthStep{}

func init() {
	for _, step := range steps {
		rawSteps[step.ID()] = step.ToProtoV1()
	}
}

func New(keyManager key.KeyManager) *AuthV1 {
	return &AuthV1{
		keyManager: keyManager,
	}
}

// Key responds with the homeserver's public key
func (v1 *AuthV1) Key(context.Context, *authv1.KeyRequest) (*authv1.KeyResponse, error) {
	return &authv1.KeyResponse{
		Key: v1.keyManager.GetPublicKey(),
	}, nil
}

func (v1 *AuthV1) NextStep(c context.Context, r *authv1.NextStepRequest) (*authv1.NextStepResponse, error) {
	return &authv1.NextStepResponse{
		Step: rawSteps["initial-step"],
	}, nil
}
