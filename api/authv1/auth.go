package authv1impl

import (
	"context"

	"github.com/harmony-development/legato/db"
	dynamicauth "github.com/harmony-development/legato/dynamic_auth"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/key"
	"github.com/thanhpk/randstr"
)

type AuthV1 struct {
	authv1.DefaultAuthService
	keyManager key.KeyManager
	db         db.AuthDB
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

func New(keyManager key.KeyManager, db db.AuthDB) *AuthV1 {
	return &AuthV1{
		keyManager: keyManager,
		db:         db,
	}
}

// Key responds with the homeserver's public key
func (v1 *AuthV1) Key(context.Context, *authv1.KeyRequest) (*authv1.KeyResponse, error) {
	return &authv1.KeyResponse{
		Key: v1.keyManager.GetPublicKey(),
	}, nil
}

func (v1 *AuthV1) BeginAuth(c context.Context, r *authv1.BeginAuthRequest) (*authv1.BeginAuthResponse, error) {
	id := randstr.Hex(16)
	if err := v1.db.SetStep(id, "initial-step"); err != nil {
		return nil, err
	}
	return &authv1.BeginAuthResponse{
		AuthId: id,
	}, nil
}

func (v1 *AuthV1) NextStep(c context.Context, r *authv1.NextStepRequest) (*authv1.NextStepResponse, error) {
	currentStep, err := v1.db.GetCurrentStep(r.AuthId)
	if err != nil {
		return nil, err
	}
	return &authv1.NextStepResponse{
		Step: rawSteps[currentStep],
	}, nil
}
