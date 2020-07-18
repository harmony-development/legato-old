package v1

import (
	"context"

	profilev1 "github.com/harmony-development/legato/gen/profile"
)

// UsernameUpdate handles the protocol's UsernameUpdate request
func (v1 *V1) UsernameUpdate(c context.Context, r *profilev1.UsernameUpdateRequest) (*profilev1.UsernameUpdateResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	panic("implement me lol")
}
