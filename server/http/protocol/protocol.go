package protocol

import (
	"time"

	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/routing"
	"github.com/harmony-development/legato/server/http/socket"
	"github.com/harmony-development/legato/server/logger"

	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
)

type API struct {
	*echo.Group
	Deps *Dependencies
}

type Dependencies struct {
	Router      *routing.Router
	APIGroup    *echo.Group
	Socket      *socket.Handler
	DB          db.IHarmonyDB
	Logger      logger.ILogger
	AuthManager *auth.Manager
	Sonyflake   *sonyflake.Sonyflake
	Config      *config.Config
}

func New(deps *Dependencies) *API {
	protocol := deps.APIGroup.Group("/protocol")
	api := &API{
		Group: protocol,
		Deps:  deps,
	}
	api.Any("/socket", func(c echo.Context) error {
		deps.Socket.Handle(c.Response(), c.Request())
		return nil
	})
	deps.Router.BindRoutes(protocol, []routing.Route{
		{
			Path:    "/connect",
			Handler: api.Connect,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    1024,
			},
			Auth:   true,
			Schema: ConnectData{},
		},
		{
			Path:    "/login",
			Handler: api.Login,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 10 * time.Second,
				Burst:    8,
			},
			Auth:   false,
			Schema: LoginData{},
		},
		{
			Path:    "/register",
			Handler: api.Register,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 15 * time.Second,
				Burst:    4,
			},
			Auth:   false,
			Schema: RegisterData{},
		},
		{
			Path:    "/key",
			Handler: api.GetKey,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    1024,
			},
			Auth:   false,
			Schema: nil,
		},
	})
	return api
}
