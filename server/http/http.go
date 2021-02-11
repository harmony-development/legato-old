package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/harmony-development/legato/server/http/attachments"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/routing"
	"github.com/harmony-development/legato/server/logger"
)

// Server is an instance of the HTTP server
type Server struct {
	Router *routing.Router
	Dependencies
}

// Dependencies are elements that a Server needs
type Dependencies struct {
	DB             types.IHarmonyDB
	Logger         logger.ILogger
	Config         *config.Config
	StorageBackend backend.AttachmentBackend
}

// New creates the /_harmony group of stuff
func New(e *echo.Echo, deps Dependencies) {
	m := hm.New(deps.DB, deps.Logger, deps.Config)
	s := &Server{
		Dependencies: deps,
		Router:       &routing.Router{Middlewares: m},
	}

	harmony := e.Group("/_harmony")
	harmony.Use(m.WithHarmony)
	harmony.Use(middleware.CORS())

	attachmentsGrp := harmony.Group("/media")
	if _, err := attachments.New(attachments.Dependencies{
		APIGroup:    attachmentsGrp,
		Router:      s.Router,
		FileBackend: s.StorageBackend,
	}); err != nil {
		panic(err)
	}
}
