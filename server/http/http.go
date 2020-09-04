package http

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sony/sonyflake"

	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/routing"
	"github.com/harmony-development/legato/server/http/webrtc"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/storage"
)

// Server is an instance of the HTTP server
type Server struct {
	*echo.Echo
	Router *routing.Router
	Dependencies
}

// Dependencies are elements that a Server needs
type Dependencies struct {
	DB             db.IHarmonyDB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Logger         logger.ILogger
	Config         *config.Config
	Sonyflake      *sonyflake.Sonyflake
}

// New creates a new HTTP server instance
func New(deps Dependencies) *Server {
	s := &Server{
		Echo:         echo.New(),
		Dependencies: deps,
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
		Validator: validator.New(),
	}
	m := hm.New(deps.DB, deps.Logger)
	s.Router = &routing.Router{Middlewares: m}
	api := s.Group("/api")
	api.Use(m.WithHarmony)

	webrtc.New(webrtc.Dependencies{})

	return s
}
