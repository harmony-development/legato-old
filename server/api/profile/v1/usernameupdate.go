package v1

import (
	profilev1 "harmony-server/gen/profile"
	"harmony-server/server/api/middleware"
)

// UsernameUpdate handles the protocol's UsernameUpdate request
func (v1 *V1) UsernameUpdate(ctx *middleware.HarmonyContext, r *profilev1.UsernameUpdateRequest) (*profilev1.UsernameUpdateResponse, error) {
	panic("implement me lol")
}
