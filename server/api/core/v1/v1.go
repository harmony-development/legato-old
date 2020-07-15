package v1

import (
	corev1 "github.com/harmony-development/legato/gen/core"
)

// V1 contains the gRPC handler for v1
type V1 struct {
	corev1.UnimplementedCoreServiceServer
}
