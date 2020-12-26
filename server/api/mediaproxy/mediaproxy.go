package mediaproxy

import (
	v1 "github.com/harmony-development/legato/server/api/mediaproxy/v1"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB     db.IHarmonyDB
	Logger logger.ILogger
	Config *config.Config
}

// Service contains the chat service
type Service struct {
	*Dependencies
	*v1.V1
}

// New creates a new chat service
func New(deps *Dependencies) *Service {
	mediaproxy := &Service{
		Dependencies: deps,
	}
	mediaproxy.V1 = &v1.V1{
		Dependencies: v1.Dependencies{
			DB:     deps.DB,
			Logger: deps.Logger,
			Config: deps.Config,
		},
	}
	mediaproxy.Initialize()
	return mediaproxy
}
