package core

import (
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	v1 "harmony-server/server/http/core/v1"

	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	"harmony-server/server/http/routing"
	"harmony-server/server/logger"
	"harmony-server/server/state"
	"harmony-server/server/storage"
)

// API contains the APIs for CoreKit
type API struct {
	*echo.Group
	Deps *Dependencies
	V1 *v1.Handlers
}

// Dependencies are items that an API needs to function
type Dependencies struct {
	Router         routing.IRouter
	APIGroup       *echo.Group
	DB             db.IHarmonyDB
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Logger         *logger.Logger
	State          *state.State
	Sonyflake      *sonyflake.Sonyflake
}

// New instantiates the handlers for CoreKit
func New(deps *Dependencies) *API {
	core := deps.APIGroup.Group("/core")
	api := &API{
		Group: core,
		Deps:  deps,
	}
	V1 := v1.New(&v1.Dependencies{
		DB:             api.Deps.DB,
		Config:         api.Deps.Config,
		AuthManager:    api.Deps.AuthManager,
		StorageManager: api.Deps.StorageManager,
		Logger:         api.Deps.Logger,
		State:          api.Deps.State,
		Sonyflake:      api.Deps.Sonyflake,
	})
	api.V1 = V1
	api.Deps.Router.BindRoutes(api.Group.Group("/v1"), V1.MakeRoutes())
	return api
}
