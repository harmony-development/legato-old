package v1

import (
	"context"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"time"
	"unicode"

	"github.com/dgrijalva/jwt-go"
	foundationv1 "github.com/harmony-development/legato/gen/foundation"
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
}

type V1 struct {
	Dependencies
}

func (v1 *V1) Federate(c context.Context, r *foundationv1.FederateRequest) (*foundationv1.FederateReply, error) {
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

	return &foundationv1.FederateReply{
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
	}, "/protocol.foundation.v1.FoundationService/Federate")
}

func (v1 *V1) Key(c context.Context, r *foundationv1.KeyRequest) (*foundationv1.KeyReply, error) {
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
	return &foundationv1.KeyReply{
		Key: string(pemData),
	}, nil
}

func (v1 *V1) LocalLogin(c context.Context, r *foundationv1.LoginRequest_Local) (*foundationv1.Session, error) {
	user, err := v1.DB.GetUserByEmail(r.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, responses.InvalidEmail)
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, r.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, responses.InvalidPassword)
	}
	session := randstr.Hex(16)
	if err := v1.DB.AddSession(user.UserID, session); err != nil {
		return nil, err
	}

	return &foundationv1.Session{
		UserId:       user.UserID,
		SessionToken: session,
	}, nil
}

func (v1 *V1) FederatedLogin(c context.Context, r *foundationv1.LoginRequest_Federated) (*foundationv1.Session, error) {
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

	return &foundationv1.Session{
		UserId:       localUserID,
		SessionToken: session,
	}, nil
}

func (v1 *V1) Login(c context.Context, r *foundationv1.LoginRequest) (*foundationv1.Session, error) {
	switch r.GetLogin().(type) {
	case *foundationv1.LoginRequest_Federated_:
		return v1.FederatedLogin(c, r.GetFederated())
	case *foundationv1.LoginRequest_Local_:
		return v1.LocalLogin(c, r.GetLocal())
	default:
		panic("invalid case")
	}
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

func (v1 *V1) Register(c context.Context, r *foundationv1.RegisterRequest) (*foundationv1.Session, error) {
	if len(r.Username) < v1.Config.Server.Policies.Username.MinLength || len(r.Username) > v1.Config.Server.Policies.Username.MaxLength {
		_ = responses.UsernameLength(
			v1.Config.Server.Policies.Username.MinLength,
			v1.Config.Server.Policies.Username.MaxLength,
		)
		return nil, status.Error(codes.InvalidArgument, responses.InvalidUsername)
	}
	if len(r.Password) < v1.Config.Server.Policies.Password.MinLength || len(r.Password) > v1.Config.Server.Policies.Password.MaxLength {
		_ = responses.PasswordLength(
			v1.Config.Server.Policies.Password.MinLength,
			v1.Config.Server.Policies.Password.MaxLength,
		)
		return nil, status.Error(codes.InvalidArgument, responses.InvalidPassword)
	}
	if !v1.PasswordAcceptable(r.Password) {
		_ = responses.PasswordPolicy(
			v1.Config.Server.Policies.Password.MinUpper,
			v1.Config.Server.Policies.Password.MinLower,
			v1.Config.Server.Policies.Password.MinNumbers,
			v1.Config.Server.Policies.Password.MinSymbols,
		)
		return nil, status.Error(codes.InvalidArgument, responses.InvalidPassword)
	}

	hash, err := bcrypt.GenerateFromPassword(r.Password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	exists, err := v1.DB.EmailExists(r.Email)
	if err != nil {
		return nil, err
	} else if exists {
		return nil, status.Error(codes.AlreadyExists, responses.AlreadyRegistered)
	}

	userID, err := v1.Sonyflake.NextID()
	if err != nil {
		return nil, err
	}

	if err := v1.DB.AddLocalUser(userID, r.Email, r.Username, hash); err != nil {
		return nil, err
	}

	session := randstr.Hex(16)
	if err := v1.DB.AddSession(userID, session); err != nil {
		return nil, err
	}

	return &foundationv1.Session{
		UserId:       userID,
		SessionToken: session,
	}, nil
}

func (v1 *V1) GetConfig(c context.Context, r *emptypb.Empty) (*foundationv1.GetConfigResponse, error) {
	return nil, nil
}
