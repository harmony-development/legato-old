// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package authv1impl

import (
	"context"

	"github.com/harmony-development/legato/api"
	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/harmony-development/legato/db/persist"
	dynamicauth "github.com/harmony-development/legato/dynamic_auth"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/key"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

type AuthV1 struct {
	authv1.DefaultAuthService
	keyManager key.KeyManager
	eph        ephemeral.Database
	persist    persist.Database
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

func New(keyManager key.KeyManager, eph ephemeral.Database, persist persist.Database) *AuthV1 {
	return &AuthV1{
		keyManager: keyManager,
		eph:        eph,
		persist:    persist,
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
	if err := v1.eph.SetStep(c, id, initialStep.ID()); err != nil {
		return nil, err
	}
	return &authv1.BeginAuthResponse{
		AuthId: id,
	}, nil
}

func (v1 *AuthV1) NextStep(c context.Context, r *authv1.NextStepRequest) (*authv1.NextStepResponse, error) {
	currentStepID, err := v1.eph.GetCurrentStep(c, r.AuthId)
	if err != nil {
		return nil, api.NewError(api.ErrorBadAuthID)
	}

	// the step the user is currently on
	step := steps[currentStepID]
	if choiceStep, ok := step.(*dynamicauth.ChoiceStep); ok {
		res, err := v1.choiceHandler(c, choiceStep, r)
		return &authv1.NextStepResponse{
			Step: res,
		}, err
	} else if formStep, ok := step.(*dynamicauth.FormStep); ok {
		res, err := v1.loginFormHandler(c, formStep, r)
		return &authv1.NextStepResponse{
			Step: res,
		}, err
	}

	return &authv1.NextStepResponse{
		Step: steps[currentStepID].ToProtoV1(),
	}, nil
}

// loginFormHandler handles the login form step
func (v1 *AuthV1) loginFormHandler(c context.Context, formStep *dynamicauth.FormStep, r *authv1.NextStepRequest) (*authv1.AuthStep, error) {
	f := r.GetForm()
	if f == nil {
		return formStep.ToProtoV1(), nil
	}
	if err := formStep.ValidateFormV1(f); err != nil {
		return nil, api.NewError(api.ErrorBadFormData)
	}

	email := f.Fields[0].GetString_()
	provided := f.Fields[1].GetBytes()

	user, local, err := v1.persist.Users().GetLocalByEmail(c, email)
	if err != nil {
		return nil, api.NewError(api.ErrorBadCredentials)
	}
	if err := bcrypt.CompareHashAndPassword(local.Password, provided); err != nil {
		// intentionally generic error to give less information to the user
		return nil, api.NewError(api.ErrorBadCredentials)
	}

	// login succeessful

	sessionID := randstr.Hex(16)

	if err := v1.persist.Sessions().Add(c, sessionID, user.ID); err != nil {
		return nil, err
	}

	if err := v1.eph.DeleteAuthID(c, r.AuthId); err != nil {
		return nil, err
	}

	return &authv1.AuthStep{
		Step: &authv1.AuthStep_Session{
			Session: &authv1.Session{
				UserId:       user.ID,
				SessionToken: sessionID,
			},
		},
	}, nil
}

// choiceHandler contains logic related to any choice usage
func (v1 *AuthV1) choiceHandler(ctx context.Context, choiceStep *dynamicauth.ChoiceStep, r *authv1.NextStepRequest) (*authv1.AuthStep, error) {
	c := r.GetChoice()
	if c == nil {
		return choiceStep.ToProtoV1(), nil
	}
	if !choiceStep.HasOption(c.Choice) {
		return nil, api.NewError(api.ErrorBadChoice)
	}

	if err := v1.eph.SetStep(ctx, r.AuthId, c.Choice); err != nil {
		return nil, api.NewError(api.ErrorInternalServerError)
	}

	nextStep := steps[c.Choice]
	return nextStep.ToProtoV1(), nil
}
