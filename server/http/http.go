package http

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/attachments"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/routing"
	"github.com/harmony-development/legato/server/http/webrtc"
	"github.com/harmony-development/legato/server/logger"
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
	Logger         logger.ILogger
	Config         *config.Config
	StorageBackend backend.AttachmentBackend
}

// New creates a new HTTP server instance
func New(deps Dependencies) *Server {
	s := &Server{
		Echo:         echo.New(),
		Dependencies: deps,
	}
	s.Pre(middleware.RemoveTrailingSlash())
	if deps.Config.Server.UseCORS {
		s.Use(middleware.CORS())
	}
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

	harmony := s.Group("/_harmony")
	harmony.Use(m.WithHarmony)

	webrtcGrp := harmony.Group("/webrtc")
	webrtc.New(webrtc.Dependencies{
		APIGroup: webrtcGrp,
		Router:   s.Router,
	})

	attachmentsGrp := harmony.Group("/media")
	if _, err := attachments.New(attachments.Dependencies{
		APIGroup:    attachmentsGrp,
		Router:      s.Router,
		FileBackend: s.StorageBackend,
	}); err != nil {
		panic(err)
	}

	return s
}
