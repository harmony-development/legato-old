package foundation

import (
	v1 "github.com/harmony-development/legato/server/api/foundation/v1"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/sony/sonyflake"
)

type Dependencies struct {
	DB          db.IHarmonyDB
	AuthManager *auth.Manager
	Sonyflake   *sonyflake.Sonyflake
	Config      *config.Config
}

type Service struct {
	*Dependencies
	*v1.V1
}

func New(deps *Dependencies) *Service {
	foundation := &Service{
		Dependencies: deps,
	}
	foundation.V1 = &v1.V1{
		Dependencies: v1.Dependencies{
			DB:          deps.DB,
			AuthManager: deps.AuthManager,
			Sonyflake:   deps.Sonyflake,
			Config:      deps.Config,
		},
	}
	return foundation
}
