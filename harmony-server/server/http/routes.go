package http

import (
	"github.com/labstack/echo/v4"
	v1 "harmony-server/server/http/v1"
)

// BindRoutes applies routes for /api
func (s Server) BindRoutes(g *echo.Group) {
	g.GET("/socket", func(ctx echo.Context) error {
		s.Socket.Handle(ctx.Response(), ctx.Request())
		return nil
	})
	s.V1 = v1.New(&v1.Dependencies{
		DB:             s.Deps.DB,
		Config:         s.Deps.Config,
		AuthManager:    s.Deps.AuthManager,
		StorageManager: s.Deps.StorageManager,
		State:          s.Deps.State,
	}, g)
}
