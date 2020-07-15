package v1

import (
	"context"
	"database/sql"
	"errors"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/http/responses"
)

// GetUser handles the protocol's GetUser request
func (v1 *V1) GetUser(c context.Context, r *profilev1.GetUserRequest) (*profilev1.GetUserResponse, error) {
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
