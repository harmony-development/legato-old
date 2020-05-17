package core

import (
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"

	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	v1 "harmony-server/server/http/core/v1"
	"harmony-server/server/http/routing"
	"harmony-server/server/state"
	"harmony-server/server/storage"
)

// API contains the APIs for CoreKit
type API struct {
	*echo.Group
	Deps *Dependencies
}

// Dependencies are items that an API needs to function
type Dependencies struct {
	Router         *routing.Router
	APIGroup       *echo.Group
	DB             *db.HarmonyDB
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	State          *state.State
	Sonyflake      *sonyflake.Sonyflake
}

// New instantiates the handlers for CoreKit
func New(deps *Dependencies) *API {
	core := deps.APIGroup.Group("/core")
	return &API{
		Group: core,
		Deps:  deps,
	}
}

// MakeRoutes creates the handlers for CoreKit
func (api *API) MakeRoutes() {
	api.Deps.Router.BindRoutes(api.Group.Group("/v1"), v1.New(&v1.Dependencies{
		DB:             api.Deps.DB,
		Config:         api.Deps.Config,
		AuthManager:    api.Deps.AuthManager,
		StorageManager: api.Deps.StorageManager,
		State:          api.Deps.State,
		Sonyflake:      api.Deps.Sonyflake,
	}).MakeRoutes())
}
