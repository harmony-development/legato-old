package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-auth-server/server/auth"
	"harmony-auth-server/server/config"
	"harmony-auth-server/server/db"
	"harmony-auth-server/server/http/hm"
	"harmony-auth-server/server/storage"
	"time"
)

// Handlers represents the handlers for API v1
type Handlers struct {
	Group          *echo.Group
	DB             *db.DB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Config         *config.Config
}

// New returns a new v1 model
func New(db *db.DB, authManager *auth.Manager, storageManager *storage.Manager, config *config.Config, g *echo.Group) *Handlers {
	h := &Handlers{
		DB:             db,
		AuthManager:    authManager,
		StorageManager: storageManager,
		Config:         config,
		Group:    g,
	}

	g.POST("/login", hm.WithRateLimit(h.Login, 5*time.Second, 5))
	g.POST("/register", hm.WithRateLimit(h.Register, 15*time.Second, 8))
	g.POST("/getuser", hm.WithRateLimit(h.GetUser, 5*time.Second, 5))
	g.POST("/usernameupdate", hm.WithRateLimit(h.UsernameUpdate, 2*time.Second, 5))
	g.POST("/auth", hm.WithRateLimit(h.Authenticate, 2*time.Second, 5))
	g.POST("/addinstance", hm.WithRateLimit(h.AddInstance, 5*time.Second, 5))
	g.POST("/removeserver", hm.WithRateLimit(h.RemoveServer, 5*time.Second, 5))
	g.POST("/listservers", hm.WithRateLimit(h.ListInstances, 5*time.Second, 5))

	return h
}
