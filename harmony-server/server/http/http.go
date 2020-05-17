package http

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sony/sonyflake"

	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	"harmony-server/server/http/core"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/routing"
	"harmony-server/server/http/socket"
	"harmony-server/server/state"
	"harmony-server/server/storage"
)

// Server is an instance of the HTTP server
type Server struct {
	*echo.Echo
	Router  *routing.Router
	CoreAPI *core.API
	Socket  *socket.Handler
	Deps    *Dependencies
}

// Dependencies are elements that a Server needs
type Dependencies struct {
	DB             *db.HarmonyDB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	State          *state.State
	Config         *config.Config
	Sonyflake      *sonyflake.Sonyflake
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
	m := hm.New(deps.DB)
	s.Router = &routing.Router{Middlewares: m}
	api := s.Group("/api")
	api.Use(m.WithHarmony)
	s.CoreAPI = core.New(&core.Dependencies{
		Router:         s.Router,
		APIGroup:       api,
		DB:             s.Deps.DB,
		Config:         s.Deps.Config,
		AuthManager:    s.Deps.AuthManager,
		StorageManager: s.Deps.StorageManager,
		State:          s.Deps.State,
		Sonyflake:      s.Deps.Sonyflake,
	})
	s.CoreAPI.MakeRoutes()
	return s
}
