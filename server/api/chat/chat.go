package chat

import (
	v1 "github.com/harmony-development/legato/server/api/chat/v1"
	"github.com/harmony-development/legato/server/api/chat/v1/permissions"
	"github.com/harmony-development/legato/server/api/chat/v1/pubsub_backends/inprocess"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"
)

// Dependencies are the backend services this package needs
type Dependencies struct {
	DB             types.IHarmonyDB
	Logger         logger.ILogger
	Sonyflake      *sonyflake.Sonyflake
	Perms          *permissions.Manager
	Config         *config.Config
	Middlewares    *middleware.Middlewares
	StorageBackend backend.AttachmentBackend
}

// Service contains the chat service
type Service struct {
	*Dependencies
	V1 *v1.V1
}

// New creates a new chat service
func New(deps *Dependencies) *Service {
	chat := &Service{
		Dependencies: deps,
	}

	pubsub := &inprocess.StreamManager{}
	pubsub.Init(deps.Logger, deps.DB)

	chat.V1 = &v1.V1{
		Dependencies: v1.Dependencies{
			DB:             deps.DB,
			Logger:         deps.Logger,
			Sonyflake:      deps.Sonyflake,
			Perms:          deps.Perms,
			Streams:        pubsub,
			Config:         deps.Config,
			StorageBackend: deps.StorageBackend,
			Middleware:     chat.Middlewares,
		},
	}
	return chat
}
