package http

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/socket"
	v1 "harmony-server/server/http/v1"
	"harmony-server/server/state"
	"harmony-server/server/storage"
)

// Server is an instance of the HTTP server
type Server struct {
	*echo.Echo
	Socket *socket.Handler
	V1     *v1.Handlers
	Deps   *Dependencies
}

type Dependencies struct {
	DB             *db.DB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	State          *state.State
	Config         *config.Config
}

// New creates a new HTTP server instance
func New(deps *Dependencies) *Server {
	s := &Server{
		Echo:   echo.New(),
		Socket: socket.NewHandler(deps.State),
		Deps:   deps,
	}
	s.Pre(middleware.RemoveTrailingSlash())
	s.Use(middleware.CORS())
	s.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:       1 << 10,
		DisableStackAll: true,
	}))
	s.Use(sentryecho.New(sentryecho.Options{
		Repanic:         true,
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
