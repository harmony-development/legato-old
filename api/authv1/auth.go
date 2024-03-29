// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package authv1impl

import (
	"context"

	"github.com/harmony-development/legato/api"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db/ephemeral"
	"github.com/harmony-development/legato/db/persist"
	dynamicauth "github.com/harmony-development/legato/dynamic_auth"
	"github.com/harmony-development/legato/errwrap"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/id"
	"github.com/harmony-development/legato/key"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

type formHandlerFunc = func(
	c context.Context,
	submission *authv1.NextStepRequest_Form,
	r *authv1.NextStepRequest,
) (
	*authv1.AuthStep,
	error,
)

type AuthV1 struct {
	authv1.DefaultAuthService
	keyManager   key.Manager
	eph          ephemeral.Database
	persist      persist.Database
	cfg          *config.Config
	idGen        id.Generator
	steps        map[string]dynamicauth.Step
	formHandlers map[string]formHandlerFunc
}

func toStepMap(steps ...dynamicauth.Step) map[string]dynamicauth.Step {
	ret := map[string]dynamicauth.Step{}
	for _, step := range steps {
		ret[step.ID()] = step
	}

	return ret
}

func New(
	keyManager key.Manager,
	eph ephemeral.Database,
	persist persist.Database,
	idGen id.Generator,
	cfg *config.Config,
) *AuthV1 {
	a := &AuthV1{
		keyManager: keyManager,
		eph:        eph,
		persist:    persist,
		cfg:        cfg,
		idGen:      idGen,
		steps: toStepMap(
			initialStep,
			loginStep,
			registerStep,
			otherOptionsStep,
			resetPasswordStep,
		),
	}

	a.formHandlers = map[string]formHandlerFunc{
		loginStep.ID():    a.loginFormHandler,
		registerStep.ID(): a.registerHandler,
	}

	return a
}

// Key responds with the homeserver's public key.
func (v1 *AuthV1) Key(context.Context, *authv1.KeyRequest) (*authv1.KeyResponse, error) {
	return &authv1.KeyResponse{
		Key: v1.keyManager.GetPublicKey(),
	}, nil
}

func (v1 *AuthV1) BeginAuth(c context.Context, r *authv1.BeginAuthRequest) (*authv1.BeginAuthResponse, error) {
	sessionID := randstr.Hex(v1.cfg.AuthIDLength)
	if err := v1.eph.SetStep(c, sessionID, initialStep.ID()); err != nil {
		return nil, errwrap.Wrap(err, "failed to set step")
	}

	return &authv1.BeginAuthResponse{
		AuthId: sessionID,
	}, nil
}

// NextStep handles dyhnamic auth steps.
func (v1 *AuthV1) NextStep(ctx context.Context, r *authv1.NextStepRequest) (*authv1.NextStepResponse, error) {
	// the ID of the step the user is on
	currentStepID, err := v1.eph.GetCurrentStep(ctx, r.AuthId)
	if err != nil {
		return nil, api.NewError(api.ErrorBadAuthID)
	}

	res, err := v1.handleStep(ctx, v1.steps[currentStepID], r)

	return &authv1.NextStepResponse{
		Step: res,
	}, err
}

func (v1 *AuthV1) handleStep(
	ctx context.Context,
	currentStep dynamicauth.Step,
	r *authv1.NextStepRequest,
) (
	*authv1.AuthStep,
	error,
) {
	switch currentStep := currentStep.(type) {
	case *dynamicauth.ChoiceStep:
		return v1.choiceHandler(ctx, currentStep, r)
	case *dynamicauth.FormStep:
		formSubmission := r.GetForm()
		if formSubmission == nil {
			return currentStep.ToProtoV1(), nil
		}

		if err := currentStep.ValidateFormV1(formSubmission); err != nil {
			return nil, api.NewError(api.ErrorBadFormData)
		}

		return v1.formHandlers[currentStep.ID()](ctx, formSubmission, r)
	default:
		return nil, errwrap.Wrap(api.NewError(api.ErrorBadStep), "invalid auth step")
	}
}

// choiceHandler contains logic related to any choice step.
func (v1 *AuthV1) choiceHandler(
	ctx context.Context,
	choiceStep *dynamicauth.ChoiceStep,
	r *authv1.NextStepRequest,
) (
	*authv1.AuthStep,
	error,
) {
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

	nextStep := v1.steps[c.Choice]

	return nextStep.ToProtoV1(), nil
}

// loginFormHandler handles the login form step.
func (v1 *AuthV1) loginFormHandler(
	c context.Context,
	submission *authv1.NextStepRequest_Form,
	r *authv1.NextStepRequest,
) (
	*authv1.AuthStep,
	error,
) {
	email := submission.Fields[0].GetString_()
	provided := submission.Fields[1].GetBytes()

	user, local, err := v1.persist.Users().GetLocalByEmail(c, email)
	if err != nil {
		return nil, api.NewError(api.ErrorBadCredentials)
	}

	if err := bcrypt.CompareHashAndPassword(local.Password, provided); err != nil {
		// intentionally generic error to give less information to the user
		return nil, api.NewError(api.ErrorBadCredentials)
	}

	// login succeessful
	session, err := v1.finishAuth(c, r.AuthId, user.ID)
	if err != nil {
		return nil, err
	}

	return &authv1.AuthStep{
		Step: &authv1.AuthStep_Session{
			Session: session,
		},
	}, nil
}

// registerHandler handles the register form step.
func (v1 *AuthV1) registerHandler(
	c context.Context,
	submission *authv1.NextStepRequest_Form,
	r *authv1.NextStepRequest,
) (
	*authv1.AuthStep,
	error,
) {
	email := submission.Fields[0].GetString_()
	username := submission.Fields[1].GetString_()
	password := submission.Fields[2].GetBytes()

	id, err := v1.idGen.NextID()
	if err != nil {
		return nil, errwrap.Wrap(err, "unable to generate new user ID")
	}

	pass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, api.NewError(api.ErrorInternalServerError)
	}

	if err := v1.persist.Users().Add(c, persist.UserInformation{
		ID:       id,
		Username: username,
	}, persist.LocalUserInformation{
		Email:    email,
		Password: pass,
	}); err != nil {
		return nil, errwrap.Wrap(err, "failed to add user")
	}

	// login succeessful
	session, err := v1.finishAuth(c, r.AuthId, id)
	if err != nil {
		return nil, err
	}

	return &authv1.AuthStep{
		Step: &authv1.AuthStep_Session{
			Session: session,
		},
	}, nil
}

func (v1 *AuthV1) finishAuth(c context.Context, authID string, userID uint64) (*authv1.Session, error) {
	sessionID := randstr.Hex(v1.cfg.AuthIDLength)

	if err := v1.persist.Sessions().Add(c, sessionID, userID); err != nil {
		return nil, errwrap.Wrap(err, "failed to add session")
	}

	if err := v1.eph.DeleteAuthID(c, authID); err != nil {
		return nil, errwrap.Wrap(err, "failed to delete auth ID")
	}

	return &authv1.Session{
		UserId:       userID,
		SessionToken: sessionID,
	}, nil
}
