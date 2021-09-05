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
	auth       db.AuthDB
}

var rawSteps = toRawSteps(
	initialStep,
	loginStep,
	registerStep,
	otherOptionsStep,
	resetPasswordStep,
)

func toRawSteps(steps ...dynamicauth.Step) map[string]*authv1.AuthStep {
	ret := map[string]*authv1.AuthStep{}
	for _, step := range steps {
		ret[step.ID()] = step.ToProtoV1()
	}
	return ret
}

func New(keyManager key.KeyManager, auth db.AuthDB) *AuthV1 {
	return &AuthV1{
		keyManager: keyManager,
		auth:       auth,
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
	// typesafe ftw
	if err := v1.auth.SetStep(id, initialStep.ID()); err != nil {
		return nil, err
	}
	return &authv1.BeginAuthResponse{
		AuthId: id,
	}, nil
}

func (v1 *AuthV1) NextStep(c context.Context, r *authv1.NextStepRequest) (*authv1.NextStepResponse, error) {
	currentStep, err := v1.auth.GetCurrentStep(r.AuthId)
	if err != nil {
		return nil, err
	}
	return &authv1.NextStepResponse{
		Step: rawSteps[currentStep],
	}, nil
}
