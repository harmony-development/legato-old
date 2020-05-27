package routing

import (
	"time"

	"github.com/labstack/echo/v4"

	"harmony-server/server/http/hm"
)

const (
	GET = iota
	POST
	PUT
	DELETE
	PATCH
	ANY
)

type LocationType int

const (
	LocationNone = iota
	LocationGuild
	LocationGuildAndChannel
	LocationGuildChannelAndMessage
)

type RateLimit struct {
	Duration time.Duration
	Burst    int
}

type Route struct {
	*echo.Group
	Path        string
	Handler     echo.HandlerFunc
	RateLimit   *RateLimit
	Auth        bool
	Schema      interface{}
	Method      int
	Location    LocationType
	Permissions hm.Permission
}

type Router struct {
	Middlewares *hm.Middlewares
}

// BindRoute binds a route to an echo.Group
func (r Router) BindRoute(g *echo.Group, endpoint Route) {
	var middleware []echo.MiddlewareFunc
	if endpoint.Auth {
		middleware = append(middleware, r.Middlewares.WithAuth)
	}
	if endpoint.RateLimit != nil {
		middleware = append(middleware, r.Middlewares.RateLimit(endpoint.RateLimit.Duration, endpoint.RateLimit.Burst))
	}
	if endpoint.Schema != nil {
		middleware = append(middleware, r.Middlewares.Schema(endpoint.Schema))
	}
	switch endpoint.Location {
	case LocationNone: // Do nothing.
	case LocationGuild:
		middleware = append(middleware, r.Middlewares.WithGuild)
	case LocationGuildAndChannel:
		middleware = append(middleware, r.Middlewares.WithGuild, r.Middlewares.WithChannel)
	case LocationGuildChannelAndMessage:
		middleware = append(middleware, r.Middlewares.WithGuild, r.Middlewares.WithChannel, r.Middlewares.WithMessage)
	}
	if endpoint.Permissions != 0 {
		switch endpoint.Location {
		case LocationGuild:
			middleware = append(middleware, r.Middlewares.ForGuildPermission(endpoint.Permissions))
		case LocationGuildAndChannel, LocationGuildChannelAndMessage:
			middleware = append(middleware, r.Middlewares.ForChannelPermission(endpoint.Permissions))
		}
	}
	switch endpoint.Method {
	case GET:
		{
			g.GET(endpoint.Path, endpoint.Handler, middleware...)
		}
	case POST:
		{
			g.POST(endpoint.Path, endpoint.Handler, middleware...)
		}
	case PUT:
		{
			g.PUT(endpoint.Path, endpoint.Handler, middleware...)
		}
	case DELETE:
		{
			g.DELETE(endpoint.Path, endpoint.Handler, middleware...)
		}
	case PATCH:
		{
			g.PATCH(endpoint.Path, endpoint.Handler, middleware...)
		}
	case ANY:
		{
			g.Any(endpoint.Path, endpoint.Handler, middleware...)
		}
	}
}

// BindRoutes binds multiple Routes
func (r Router) BindRoutes(g *echo.Group, endPoints []Route) {
	for _, endPoint := range endPoints {
		r.BindRoute(g, endPoint)
	}
}
