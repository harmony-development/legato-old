package v1

import (
	profilev1 "harmony-server/gen/profile"
	"harmony-server/server/db"
	"harmony-server/server/logger"
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
