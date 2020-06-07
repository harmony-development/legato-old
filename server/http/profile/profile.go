package profile

import (
	"harmony-server/server/config"
	"harmony-server/server/db"
	v1 "harmony-server/server/http/profile/v1"
	"harmony-server/server/http/routing"
	"harmony-server/server/logger"
	"harmony-server/server/state"
	"harmony-server/server/storage"

	"github.com/labstack/echo/v4"
)

// API contains the APIs for ProfileKit
type API struct {
	*echo.Group
	Deps *Dependencies
}

// Dependencies are the elements that ProfileKit handlers need
type Dependencies struct {
	Router         *routing.Router
	APIGroup       *echo.Group
	DB             db.IHarmonyDB
	Config         *config.Config
	StorageManager *storage.Manager
	Logger         *logger.Logger
	State          *state.State
}

// New instantiates the handlers for ProfileKit
func New(deps *Dependencies) *API {
	profile := deps.APIGroup.Group("/profile")
	api := &API{
		Group: profile,
		Deps:  deps,
	}
	api.Deps.Router.BindRoutes(api.Group.Group("/v1"), v1.New(&v1.Dependencies{
		DB:             deps.DB,
		Config:         deps.Config,
		StorageManager: deps.StorageManager,
		Logger:         deps.Logger,
		State:          deps.State,
	}).MakeRoutes())
	return api
}
