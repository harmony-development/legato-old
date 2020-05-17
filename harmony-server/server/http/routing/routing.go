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

type RateLimit struct {
	Duration time.Duration
	Burst    int
}

type Route struct {
	*echo.Group
	Path      string
	Handler   echo.HandlerFunc
	RateLimit *RateLimit
	Auth      bool
	Schema    interface{}
	Method    int
}

type Router struct {
	Middlewares *hm.Middlewares
}

// BindRoute binds a route to an echo.Group
func (r Router) BindRoute(g *echo.Group, endPoint Route) {
	var middleware []echo.MiddlewareFunc
	if endPoint.Auth {
		middleware = append(middleware, r.Middlewares.WithAuth)
	}
	if endPoint.RateLimit != nil {
		middleware = append(middleware, r.Middlewares.RateLimit(endPoint.RateLimit.Duration, endPoint.RateLimit.Burst))
	}
	if endPoint.Schema != nil {
		middleware = append(middleware, r.Middlewares.Schema(endPoint.Schema))
	}
	switch endPoint.Method {
	case GET:
		{
			g.GET(endPoint.Path, endPoint.Handler, middleware...)
		}
	case POST:
		{
			g.POST(endPoint.Path, endPoint.Handler, middleware...)
		}
	case PUT:
		{
			g.PUT(endPoint.Path, endPoint.Handler, middleware...)
		}
	case DELETE:
		{
			g.DELETE(endPoint.Path, endPoint.Handler, middleware...)
		}
	case PATCH:
		{
			g.PATCH(endPoint.Path, endPoint.Handler, middleware...)
		}
	case ANY:
		{
			g.Any(endPoint.Path, endPoint.Handler, middleware...)
		}
	}
}

// BindRoutes binds multiple Routes
func (r Router) BindRoutes(g *echo.Group, endPoints []Route) {
	for _, endPoint := range endPoints {
		r.BindRoute(g, endPoint)
	}
}
