package authv1impl

import (
	"context"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/server/key"
)

type AuthV1 struct {
	authv1.DefaultAuthService
	keyManager key.KeyManager
}

func New(keyManager key.KeyManager) *AuthV1 {
	return &AuthV1{
		keyManager: keyManager,
	}
}

func (v1 *AuthV1) Key(context.Context, *authv1.KeyRequest) (*authv1.KeyResponse, error) {
	return &authv1.KeyResponse{
		Key: v1.keyManager.GetPublicKey(),
	}, nil
}
