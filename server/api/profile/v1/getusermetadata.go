package v1

import (
	"context"
	"database/sql"
	"errors"

	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/http/responses"
)

// GetUserMetadata handles the protocol's GetUserMetadata request
func (v1 *V1) GetUserMetadata(ctx context.Context, r *profilev1.GetUserMetadataRequest) (*profilev1.GetUserMetadataResponse, error) {
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
