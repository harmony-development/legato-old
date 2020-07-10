package v1

import (
	"errors"
	profilev1 "harmony-server/gen/profile"
	"harmony-server/server/api/middleware"
	"harmony-server/server/http/responses"
)

// StatusUpdate handles the protocol's StatusUpdate request
func (v1 *V1) StatusUpdate(ctx *middleware.HarmonyContext, r *profilev1.StatusUpdateRequest) (*profilev1.StatusUpdateResponse, error) {
	if err := v1.DB.SetStatus(ctx.UserID, r.NewStatus); err != nil {
		v1.Logger.Exception(err)
		return nil, errors.New(responses.UnknownError)
	}
	return &profilev1.StatusUpdateResponse{
		Success: true,
	}, nil
}
