package v1

import (
	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
)

// Dependencies contains the services that the profile service uses
type Dependencies struct {
	DB     db.IHarmonyDB
	Logger logger.ILogger
}

// V1 contains the gRPC handler for v1
type V1 struct {
	profilev1.UnimplementedProfileServiceServer
	*Dependencies
}
