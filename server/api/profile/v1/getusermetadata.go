package v1

import (
	"context"
	"database/sql"
	"errors"
	"time"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/http/responses"
)

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 1 * time.Second,
			Burst:    4,
		},
		Auth:       true,
		Permission: middleware.NoPermission,
	}, "/protocol.profile.v1.ProfileService/GetUserMetadata")
}

// GetUserMetadata handles the protocol's GetUserMetadata request
func (v1 *V1) GetUserMetadata(ctx context.Context, r *profilev1.GetUserMetadataRequest) (*profilev1.GetUserMetadataResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	meta, err := v1.DB.GetUserMetadata(0, r.AppId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(responses.MetadataNotFound)
		}
		v1.Logger.Exception(err)
		return nil, errors.New(responses.UnknownError)
	}
	return &profilev1.GetUserMetadataResponse{
		Metadata: meta,
	}, nil
}
