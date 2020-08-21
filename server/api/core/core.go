package core

import (
	v1 "github.com/harmony-development/legato/server/api/core/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB          db.IHarmonyDB
	Middlewares middleware.Middlewares
	Logger      logger.ILogger
	Sonyflake   *sonyflake.Sonyflake
}

// Service contains the core service
type Service struct {
	*Dependencies
	V1 *v1.V1
}

// New creates a new core service
func New(deps *Dependencies) *Service {
	core := &Service{
		Dependencies: deps,
	}
	core.V1 = &v1.V1{
		Dependencies: v1.Dependencies{
			DB:          deps.DB,
			Middlewares: deps.Middlewares,
			Logger:      deps.Logger,
			Sonyflake:   deps.Sonyflake,
		},
	}
	return core
}
