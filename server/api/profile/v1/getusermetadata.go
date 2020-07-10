package v1

import (
	"database/sql"
	"errors"
	profilev1 "harmony-server/gen/profile"
	"harmony-server/server/api/middleware"
	"harmony-server/server/http/responses"
)

// GetUserMetadata handles the protocol's GetUserMetadata request
func (v1 *V1) GetUserMetadata(ctx *middleware.HarmonyContext, r *profilev1.GetUserMetadataRequest) (*profilev1.GetUserMetadataResponse, error) {
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
