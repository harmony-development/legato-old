package authsvc

import (
	v1 "github.com/harmony-development/legato/server/api/authsvc/v1"
	authstate "github.com/harmony-development/legato/server/api/authsvc/v1/pubsub_backends/integrated"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
	"github.com/sony/sonyflake"
)

type Dependencies struct {
	DB          db.IHarmonyDB
	Logger      logger.ILogger
	AuthManager *auth.Manager
	Sonyflake   *sonyflake.Sonyflake
	Config      *config.Config
}

type Service struct {
	*Dependencies
	*v1.V1
}

func New(deps *Dependencies) *Service {
	svc := &Service{
		Dependencies: deps,
	}
	svc.V1 = v1.New(v1.Dependencies{
		DB:          deps.DB,
		Logger:      deps.Logger,
		AuthManager: deps.AuthManager,
		AuthState:   authstate.New(deps.Logger),
		Sonyflake:   deps.Sonyflake,
		Config:      deps.Config,
	})
	return svc
}
