package http

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"harmony-auth-server/consts"
	"harmony-auth-server/harmony/auth"
	"harmony-auth-server/harmony/config"
	"harmony-auth-server/harmony/db"
	"harmony-auth-server/harmony/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "harmony-auth-server/harmony/http/v1"
)

// Server is a wrapper for labstack echo that implements some more methods
type Server struct {
	*echo.Echo
	V1             *v1.Handlers
	DB             *db.DB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Config         *config.Config
	Consts         *consts.Constants
}

// New returns a new http server with echo
func New(db *db.DB, authManager *auth.Manager, storageManager *storage.Manager, config *config.Config, consts *consts.Constants) *Server {
	s := &Server{
		Echo:           echo.New(),
		DB:             db,
		AuthManager:    authManager,
		StorageManager: storageManager,
		Config:         config,
		Consts:         consts,
	}
	s.Pre(middleware.RemoveTrailingSlash())
	s.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:       1 << 10,
		DisableStackAll: true,
	}))
	s.Validator = &HarmonyValidator{
		validator: validator.New(),
	}
	g := s.Group("/api")
	s.BindRoutes(g)
	return s
}

// Stop gracefully shuts down the echo server
func (s Server) Stop() {
	if err := s.Shutdown(context.TODO()); err != nil {
		logrus.Error("Error shutting down echo", err)
	}
}
