package http

import (
	"context"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"harmony-auth-server/server/auth"
	"harmony-auth-server/server/config"
	"harmony-auth-server/server/db"
	"harmony-auth-server/server/http/hm"
	"harmony-auth-server/server/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "harmony-auth-server/server/http/v1"
)

// Server is a wrapper for labstack echo that implements some more methods
type Server struct {
	*echo.Echo
	V1             *v1.Handlers
	DB             *db.DB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Config         *config.Config
}

// New returns a new http server with echo
func New(db *db.DB, authManager *auth.Manager, storageManager *storage.Manager, config *config.Config) *Server {
	s := &Server{
		Echo:           echo.New(),
		DB:             db,
		AuthManager:    authManager,
		StorageManager: storageManager,
		Config:         config,
	}
	s.Pre(middleware.RemoveTrailingSlash())
	s.Use(middleware.CORS())
	s.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:       1 << 10,
		DisableStackAll: true,
	}))
	s.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
		WaitForDelivery: false,
	}))
	s.Validator = &HarmonyValidator{
		validator: validator.New(),
	}
	g := s.Group("/api")
	s.BindRoutes(g)

	go hm.CleanupRoutine()

	return s
}

// Stop gracefully shuts down the echo server
func (s Server) Stop() {
	if err := s.Shutdown(context.TODO()); err != nil {
		logrus.Error("Error shutting down echo", err)
	}
}
