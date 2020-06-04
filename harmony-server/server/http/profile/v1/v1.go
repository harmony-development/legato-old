package v1

import (
	"time"

	"github.com/sony/sonyflake"

	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	"harmony-server/server/http/routing"
	"harmony-server/server/logger"
	"harmony-server/server/state"
	"harmony-server/server/storage"
)

// Handlers for ProfileKit
type Handlers struct {
	Deps *Dependencies
}

// Dependencies are the elements that ProfileKit handlers need
type Dependencies struct {
	DB             *db.HarmonyDB
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Logger         *logger.Logger
	State          *state.State
	Sonyflake      *sonyflake.Sonyflake
}

// New creates a new set of Handlers
func New(deps *Dependencies) *Handlers {
	return &Handlers{
		Deps: deps,
	}
}

// MakeRoutes creates the routes for ProfileKit
func (h Handlers) MakeRoutes() []routing.Route {
	return []routing.Route{
		{
			Path:    "/users/:user_id",
			Handler: h.GetUser,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    2,
			},
			Auth:     true,
			Location: routing.LocationUser,
		},
	}
}