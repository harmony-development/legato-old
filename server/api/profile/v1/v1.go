package v1

import (
	profilev1 "harmony-server/gen/profile"
)

// V1 contains the gRPC handler for v1
type V1 struct {
	profilev1.UnimplementedProfileServiceServer
}
