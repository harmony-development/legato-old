package http

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/server/http/hm"
	v1 "harmony-auth-server/server/http/v1"
)

// BindRoutes applies routes for /api
func (s Server) BindRoutes(g *echo.Group) {
	m := &hm.Middlewares{AuthManager: s.AuthManager}

	s.V1 = v1.New(&v1.Dependencies{
		DB:             s.DB,
		AuthManager:    s.AuthManager,
		StorageManager: s.StorageManager,
		Config:         s.Config,
		Middleware: m,
	}, g)
}
