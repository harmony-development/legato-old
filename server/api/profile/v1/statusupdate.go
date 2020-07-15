package v1

import (
	"context"
	"errors"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/http/responses"
)

// StatusUpdate handles the protocol's StatusUpdate request
func (v1 *V1) StatusUpdate(c context.Context, r *profilev1.StatusUpdateRequest) (*profilev1.StatusUpdateResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	if err := v1.DB.SetStatus(ctx.UserID, r.NewStatus); err != nil {
		v1.Logger.Exception(err)
		return nil, errors.New(responses.UnknownError)
	}
	return &profilev1.StatusUpdateResponse{
		Success: true,
	}, nil
}
