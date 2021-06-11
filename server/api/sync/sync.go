package sync

import (
	"github.com/enriquebris/goconcurrentqueue"
	syncv1 "github.com/harmony-development/legato/gen/sync/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	v1 "github.com/harmony-development/legato/server/api/sync/v1"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB              types.IHarmonyDB
	Logger          logger.ILogger
	Sonyflake       *sonyflake.Sonyflake
	Config          *config.Config
	AuthManager     *auth.Manager
	Middlewares     *middleware.Middlewares
	EventDispatcher func(string, *syncv1.Event)
}

// Service contains the chat service
type Service struct {
	*Dependencies
	V1 *v1.V1
}

// New creates a new chat service
func New(deps *Dependencies) *Service {
	sync := &Service{
		Dependencies: deps,
	}

	sync.V1 = &v1.V1{
		Queues: make(map[string]*goconcurrentqueue.FIFO),
		Dependencies: v1.Dependencies{
			DB:              deps.DB,
			Logger:          deps.Logger,
			Sonyflake:       deps.Sonyflake,
			Config:          deps.Config,
			AuthManager:     deps.AuthManager,
			Middleware:      sync.Middlewares,
			EventDispatcher: deps.EventDispatcher,
		},
	}
	return sync
}
