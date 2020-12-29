package v1

import (
	"context"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"time"
	"unicode"

	"github.com/dgrijalva/jwt-go"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/server/api/authsvc/v1/authsteps"
	authstate "github.com/harmony-development/legato/server/api/authsvc/v1/pubsub_backends/integrated"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/responses"
	"github.com/sony/sonyflake"
	"github.com/thanhpk/randstr"
	"github.com/ztrue/tracerr"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Dependencies struct {
	DB          db.IHarmonyDB
	Logger      logger.ILogger
	AuthManager *auth.Manager
	Sonyflake   *sonyflake.Sonyflake
	Config      *config.Config
	AuthState   *authstate.AuthState
}

type V1 struct {
	Dependencies
}

var loginStep = authsteps.NewFormStep(
	"login",
	true,
	[]authsteps.FormField{
		{
			Name:      "email",
			FieldType: "email",
		},
		{
			Name:      "password",
			FieldType: "password",
		},
	},
	[]authsteps.Step{},
)

var registerStep = authsteps.NewFormStep(
	"register",
	true,
	[]authsteps.FormField{
		{
			Name:      "email",
			FieldType: "email",
		},
		{
			Name:      "username",
			FieldType: "username",
		},
		{
			Name:      "password",
			FieldType: "password",
		},
		{
			Name:      "confirm-password",
			FieldType: "password",
		},
	},
	[]authsteps.Step{},
)

var resetPasswordStep = authsteps.NewFormStep(
	"reset-password",
	true,
	[]authsteps.FormField{
		{
			Name:      "email",
			FieldType: "email",
		},
	},
	[]authsteps.Step{},
)

var otherStep = authsteps.NewChoiceStep(
	"other-options",
	true,
	[]authsteps.Step{
		resetPasswordStep,
	},
)

var initialStep = authsteps.NewChoiceStep(
	"initial-choice",
	false,
	[]authsteps.Step{
		loginStep,
		registerStep,
		otherStep,
	},
)

func ToAuthStep(s authsteps.Step) *authv1.AuthStep {
	switch s.StepType() {
	case authsteps.StepChoice:
		{
			cs := s.(authsteps.ChoiceStep)
			return &authv1.AuthStep{
				Step: &authv1.AuthStep_Choice_{
					Choice: &authv1.AuthStep_Choice{
						Title:   cs.ID(),
						Options: cs.Choices,
					},
				},
			}
		}
	case authsteps.StepForm:
		{
			fs := s.(authsteps.FormStep)
			return &authv1.AuthStep{
				Step: &authv1.AuthStep_Form_{
					Form: &authv1.AuthStep_Form{
						Title: fs.ID(),
						Fields: func() []*authv1.AuthStep_Form_FormField {
							fields := []*authv1.AuthStep_Form_FormField{}

							for _, f := range fs.Fields {
								fields = append(fields, &authv1.AuthStep_Form_FormField{
									Name: f.Name,
									Type: f.FieldType,
								})
							}

							return fields
						}(),
					},
				},
			}
		}
	default:
		return nil
	}
}

func New(deps Dependencies) *V1 {
	if deps.Config.Server.Policies.EnablePasswordResetForm {
		otherStep.AddStep(resetPasswordStep)
	}

	return &V1{
		Dependencies: deps,
	}
}

func (v1 *V1) Federate(c context.Context, r *authv1.FederateRequest) (*authv1.FederateReply, error) {
	ctx := c.(middleware.HarmonyContext)

	user, err := v1.DB.GetUserByID(ctx.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, v1.Logger.ErrorResponse(codes.NotFound, err, "user not found")
		}
		return nil, err
	}

	token, err := v1.AuthManager.MakeAuthToken(ctx.UserID, r.Target, user.Username, user.Avatar.String)
	if err != nil {
		return nil, err
	}

	nonce := randstr.Base64(v1.Config.Server.Policies.Federation.NonceLength)
	err = v1.DB.AddNonce(nonce, user.UserID, r.Target)
	err = tracerr.Wrap(err)

	return &authv1.FederateReply{
		Token: token,
		Nonce: nonce,
	}, err
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 1 * time.Second,
			Burst:    3,
		},
		Auth: true,
	}, "/protocol.auth.v1.AuthService/Federate")
}

func (v1 *V1) Key(c context.Context, r *emptypb.Empty) (*authv1.KeyReply, error) {
	keyBytes, err := x509.MarshalPKIXPublicKey(v1.AuthManager.PubKey)
	if err != nil {
		return nil, err
	}
	pemData := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: keyBytes,
		},
	)
	return &authv1.KeyReply{
		Key: string(pemData),
	}, nil
}

func (v1 *V1) LoginFederated(c context.Context, r *authv1.LoginFederatedRequest) (*authv1.Session, error) {
	pem, err := v1.AuthManager.GetPublicKey(r.Domain)
	if err != nil {
		return nil, err
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
	if err != nil {
		return nil, err
	}

	t, err := jwt.ParseWithClaims(r.AuthToken, &auth.Token{}, func(_ *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	if err != nil {
		return nil, err
	}

	token := t.Claims.(*auth.Token)
	session := randstr.Hex(16)
	localUserID, err := v1.DB.GetLocalUserForForeignUser(token.UserID, r.Domain)

	if err != nil || localUserID == 0 {
		id, err := v1.Sonyflake.NextID()
		if err != nil {
			return nil, err
		}
		localUserID, err = v1.DB.AddForeignUser(r.Domain, token.UserID, id, token.Username, token.Avatar)
		if err != nil {
			return nil, err
		}
	}

	if err := v1.DB.AddSession(localUserID, session); err != nil {
		return nil, err
	}

	return &authv1.Session{
		UserId:       localUserID,
		SessionToken: session,
	}, nil
}

func (v1 *V1) PasswordAcceptable(passwd []byte) bool {
	var stats struct {
		upper   int
		lower   int
		numbers int
		symbols int
	}
	for _, c := range passwd {
		if unicode.IsUpper(rune(c)) {
			stats.upper++
		} else if unicode.IsLower(rune(c)) {
			stats.lower++
		} else if unicode.IsNumber(rune(c)) {
			stats.numbers++
		} else if unicode.IsSymbol(rune(c)) {
			stats.symbols++
		}
	}
	bad := stats.upper < v1.Config.Server.Policies.Password.MinUpper ||
		stats.lower < v1.Config.Server.Policies.Password.MinLower ||
		stats.numbers < v1.Config.Server.Policies.Password.MinNumbers ||
		stats.symbols < v1.Config.Server.Policies.Password.MinSymbols
	return !bad
}

func (v1 *V1) GetConfig(c context.Context, r *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (v1 *V1) BeginAuth(c context.Context, r *emptypb.Empty) (*authv1.BeginAuthResponse, error) {
	authID := randstr.Hex(32)

	if err := v1.AuthState.NewAuthSession(authID, initialStep); err != nil {
		return nil, err
	}

	go func() {
		time.Sleep(30 * time.Second)
		if !v1.AuthState.HasStream(authID) {
			v1.AuthState.DeleteAuthSession(authID)
		}
	}()

	return &authv1.BeginAuthResponse{
		AuthId: authID,
	}, nil
}

func (v1 *V1) StreamSteps(r *authv1.StreamStepsRequest, s authv1.AuthService_StreamStepsServer) error {
	channel, err := v1.AuthState.Subscribe(r.AuthId, s)
	if err != nil {
		return err
	}
	<-channel
	return nil
}

func (v1 *V1) NextStep(c context.Context, r *authv1.NextStepRequest) (*authv1.AuthStep, error) {
	if ok := v1.AuthState.AuthSessionExists(r.AuthId); !ok {
		return nil, errors.New("missing auth ID")
	}

	currentStep := v1.AuthState.GetStep(r.AuthId)

	if currentStep.StepType() == authsteps.StepChoice {
		return v1.ChoiceHandler(r)
	}

	switch currentStep.ID() {
	case "login":
		{
			return v1.LocalLogin(r)
		}
	case "register":
		{
			return v1.Register(r)
		}
	}

	return nil, nil
}

func (v1 *V1) StepBack(c context.Context, r *authv1.StepBackRequest) (*authv1.AuthStep, error) {
	if ok := v1.AuthState.AuthSessionExists(r.AuthId); !ok {
		return nil, errors.New("missing auth ID")
	}

	currentStep := v1.AuthState.GetStep(r.AuthId)

	if currentStep.CanGoBack() {
		previousStep := currentStep.GetPreviousStep()
		conv := ToAuthStep(previousStep)
		v1.AuthState.Broadcast(r.AuthId, conv)
		v1.AuthState.SetStep(r.AuthId, previousStep)
		return conv, nil
	}

	return nil, errors.New("can't go back")
}

func (v1 *V1) ChoiceHandler(r *authv1.NextStepRequest) (*authv1.AuthStep, error) {
	c := r.GetChoice()
	currentStep := v1.AuthState.GetStep(r.AuthId)

	if currentStep == nil {
		return nil, errors.New("invalid auth id")
	}

	if c == nil {
		s := ToAuthStep(v1.AuthState.GetStep(r.AuthId))
		v1.AuthState.Broadcast(r.AuthId, s)
		return s, nil
	}

	selected := c.Choice

	for _, s := range currentStep.SubSteps() {
		if s.ID() == selected {
			s.SetPreviousStep(currentStep)
			conv := ToAuthStep(s)
			v1.AuthState.Broadcast(r.AuthId, conv)
			v1.AuthState.SetStep(r.AuthId, s)
			return conv, nil
		}
	}
	return nil, errors.New("unknown choice")
}

func (v1 *V1) LocalLogin(r *authv1.NextStepRequest) (*authv1.AuthStep, error) {
	f := r.GetForm()
	if f == nil {
		return nil, errors.New("missing form")
	}

	email := f.Fields[0].GetString_()
	password := f.Fields[1].GetBytes()

	user, err := v1.DB.GetUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, responses.InvalidEmail)
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, password); err != nil {
		return nil, status.Error(codes.Unauthenticated, responses.InvalidPassword)
	}
	session := randstr.Hex(16)
	if err := v1.DB.AddSession(user.UserID, session); err != nil {
		return nil, err
	}

	s := &authv1.AuthStep{
		Step: &authv1.AuthStep_Session{
			Session: &authv1.Session{
				UserId:       user.UserID,
				SessionToken: session,
			},
		},
	}

	v1.AuthState.Broadcast(r.AuthId, s)

	defer v1.AuthState.DeleteAuthSession(r.AuthId)

	return s, nil
}

func (v1 *V1) Register(r *authv1.NextStepRequest) (*authv1.AuthStep, error) {
	f := r.GetForm()
	if f != nil {
		return nil, errors.New("missing form")
	}

	email := f.Fields[0].GetString_()
	username := f.Fields[1].GetString_()
	password := f.Fields[2].GetBytes()
	_ = f.Fields[3].GetBytes() // confirmPassword

	if len(username) < v1.Config.Server.Policies.Username.MinLength || len(username) > v1.Config.Server.Policies.Username.MaxLength {
		_ = responses.UsernameLength(
			v1.Config.Server.Policies.Username.MinLength,
			v1.Config.Server.Policies.Username.MaxLength,
		)
		return nil, status.Error(codes.InvalidArgument, responses.InvalidUsername)
	}
	if len(password) < v1.Config.Server.Policies.Password.MinLength || len(password) > v1.Config.Server.Policies.Password.MaxLength {
		_ = responses.PasswordLength(
			v1.Config.Server.Policies.Password.MinLength,
			v1.Config.Server.Policies.Password.MaxLength,
		)
		return nil, status.Error(codes.InvalidArgument, responses.InvalidPassword)
	}
	if !v1.PasswordAcceptable(password) {
		_ = responses.PasswordPolicy(
			v1.Config.Server.Policies.Password.MinUpper,
			v1.Config.Server.Policies.Password.MinLower,
			v1.Config.Server.Policies.Password.MinNumbers,
			v1.Config.Server.Policies.Password.MinSymbols,
		)
		return nil, status.Error(codes.InvalidArgument, responses.InvalidPassword)
	}

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	exists, err := v1.DB.EmailExists(email)
	if err != nil {
		return nil, err
	} else if exists {
		return nil, status.Error(codes.AlreadyExists, responses.AlreadyRegistered)
	}

	userID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}

	if err := v1.DB.AddLocalUser(userID, email, username, hash); err != nil {
		return nil, err
	}

	session := randstr.Hex(16)
	if err := v1.DB.AddSession(userID, session); err != nil {
		return nil, err
	}

	s := &authv1.AuthStep{
		Step: &authv1.AuthStep_Session{
			Session: &authv1.Session{
				UserId:       userID,
				SessionToken: session,
			},
		},
	}

	v1.AuthState.Broadcast(r.AuthId, s)

	defer v1.AuthState.DeleteAuthSession(r.AuthId)

	return s, nil
}
