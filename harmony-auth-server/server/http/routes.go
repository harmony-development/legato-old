package http

import (
	"github.com/labstack/echo/v4"
	v1 "harmony-auth-server/server/http/v1"
)

// BindRoutes applies routes for /api
func (s Server) BindRoutes(g *echo.Group) {
	s.V1 = v1.New(s.DB, s.AuthManager, s.StorageManager, s.Config, g)
}
