package v1

import (
	"database/sql"
	"errors"
	profilev1 "harmony-server/gen/profile"
	"harmony-server/server/api/middleware"
	"harmony-server/server/http/responses"
)

// GetUser handles the protocol's GetUser request
func (v1 *V1) GetUser(ctx *middleware.HarmonyContext, r *profilev1.GetUserRequest) (*profilev1.GetUserResponse, error) {
	res, err := v1.DB.GetUserByID(r.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(responses.UserNotFound)
		}
		v1.Logger.Exception(err)
		return nil, errors.New(responses.UnknownError)
	}
	return &profilev1.GetUserResponse{
		UserName:   res.Username,
		UserAvatar: res.Avatar.String,
		UserStatus: profilev1.UserStatus(res.Status),
	}, nil
}
