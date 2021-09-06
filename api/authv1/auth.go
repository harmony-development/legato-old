// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package authv1impl

import (
	"context"

	"github.com/harmony-development/legato/api"
	"github.com/harmony-development/legato/db/ephemeral"
	dynamicauth "github.com/harmony-development/legato/dynamic_auth"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/key"
	"github.com/thanhpk/randstr"
)

type AuthV1 struct {
	authv1.DefaultAuthService
	keyManager key.KeyManager
	auth       ephemeral.Database
}

var steps = toStepMap(
	initialStep,
	loginStep,
	registerStep,
	otherOptionsStep,
	resetPasswordStep,
)

func toStepMap(steps ...dynamicauth.Step) map[string]dynamicauth.Step {
	ret := map[string]dynamicauth.Step{}
	for _, step := range steps {
		ret[step.ID()] = step
	}
	return ret
}

func New(keyManager key.KeyManager, auth ephemeral.Database) *AuthV1 {
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
	if err := v1.auth.SetStep(c, id, initialStep.ID()); err != nil {
		return nil, err
	}
	return &authv1.BeginAuthResponse{
		AuthId: id,
	}, nil
}

func (v1 *AuthV1) NextStep(c context.Context, r *authv1.NextStepRequest) (*authv1.NextStepResponse, error) {
	currentStepID, err := v1.auth.GetCurrentStep(c, r.AuthId)
	if err != nil {
		return nil, api.NewError(api.ErrorBadAuthID)
	}

	step := steps[currentStepID]
	if choiceStep, ok := step.(*dynamicauth.ChoiceStep); ok {
		res, err := v1.choiceHandler(choiceStep, r)
		return &authv1.NextStepResponse{
			Step: res,
		}, err
	}

	return &authv1.NextStepResponse{
		Step: steps[currentStepID].ToProtoV1(),
	}, nil
}

func (v1 *AuthV1) choiceHandler(choiceStep *dynamicauth.ChoiceStep, r *authv1.NextStepRequest) (*authv1.AuthStep, error) {
	c := r.GetChoice()
	if c == nil {
		return choiceStep.ToProtoV1(), nil
	}
	if !choiceStep.HasOption(c.Choice) {
		return nil, api.NewError(api.ErrorBadChoice)
	}
	nextStep := steps[c.Choice]
	return nextStep.ToProtoV1(), nil
}
