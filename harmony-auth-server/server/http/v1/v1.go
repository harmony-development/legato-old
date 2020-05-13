package v1

import (
	"time"

	"github.com/labstack/echo/v4"

	"harmony-auth-server/server/auth"
	"harmony-auth-server/server/config"
	"harmony-auth-server/server/db"
	"harmony-auth-server/server/http/hm"
	"harmony-auth-server/server/storage"
)

// Handlers represents the events for API v1
type Handlers struct {
	Group          *echo.Group
	DB             *db.DB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Config         *config.Config
}

// Dependencies contains the services needed to v1
type Dependencies struct {
	DB             *db.DB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Config         *config.Config
	Middleware     *hm.Middlewares
}

// New returns a new v1 model
func New(deps *Dependencies, g *echo.Group) *Handlers {
	v1 := g.Group("/v1")
	h := &Handlers{
		DB:             deps.DB,
		AuthManager:    deps.AuthManager,
		StorageManager: deps.StorageManager,
		Config:         deps.Config,
		Group:          v1,
	}
	v1.Use(hm.WithHarmony)

	v1.POST("/login", hm.WithRateLimit(h.Login, 5*time.Second, 5))
	v1.POST("/register", hm.WithRateLimit(h.Register, 15*time.Second, 8))

	a := v1.Group("")
	a.Use(deps.Middleware.WithAuth)
	a.POST("/getuser", hm.WithRateLimit(h.GetUser, 5*time.Second, 5))
	a.POST("/usernameupdate", hm.WithRateLimit(h.UsernameUpdate, 2*time.Second, 5))
	a.POST("/auth", hm.WithRateLimit(h.Authenticate, 2*time.Second, 5))
	a.POST("/addinstance", hm.WithRateLimit(h.AddInstance, 5*time.Second, 5))
	a.POST("/removeserver", hm.WithRateLimit(h.RemoveServer, 5*time.Second, 5))
	a.POST("/listservers", hm.WithRateLimit(h.ListInstances, 5*time.Second, 5))

	return h
}
