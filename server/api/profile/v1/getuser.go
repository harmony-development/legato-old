package v1

import (
	"context"
	"database/sql"
	"time"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 10 * time.Second,
			Burst:    64,
		},
		Auth:       true,
		Permission: middleware.NoPermission,
	}, "/protocol.profile.v1.ProfileService/GetUser")
}

// GetUser handles the protocol's GetUser request
func (v1 *V1) GetUser(c context.Context, r *profilev1.GetUserRequest) (*profilev1.GetUserResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	res, err := v1.DB.GetUserByID(r.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, responses.UserNotFound)
		}
		v1.Logger.Exception(err)
		return nil, status.Error(codes.Internal, responses.UnknownError)
	}
	return &profilev1.GetUserResponse{
		UserName:   res.Username,
		UserAvatar: res.Avatar.String,
		UserStatus: profilev1.UserStatus(res.Status),
	}, nil
}
