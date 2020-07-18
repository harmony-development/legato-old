package v1

import (
	"context"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/http/responses"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UsernameUpdate handles the protocol's UsernameUpdate request
func (v1 *V1) UsernameUpdate(c context.Context, r *profilev1.UsernameUpdateRequest) (*emptypb.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if err := r.Validate(); err != nil {
		return nil, err
	}
	if err := v1.DB.UpdateUsername(ctx.UserID, r.UserName); err != nil {
		return nil, status.Error(codes.Internal, responses.UnknownError)
	}
	return &emptypb.Empty{}, nil
}
