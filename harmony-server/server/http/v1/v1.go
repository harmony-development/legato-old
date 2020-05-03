package v1

import (
	"github.com/labstack/echo/v4"
	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	"harmony-server/server/http/hm"
	"harmony-server/server/state"
	"harmony-server/server/storage"
	"time"
)

// Handlers represents the events for API v1
type Handlers struct {
	*echo.Group
	Deps *Dependencies
}

type Dependencies struct {
	DB             *db.DB
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	State          *state.State
}

// New creates a new Handlers instance
func New(deps *Dependencies, g *echo.Group) *Handlers {
	h := &Handlers{
		Group: g,
		Deps:  deps,
	}
	m := &hm.Middlewares{
		DB: deps.DB,
	}
	h.Use(m.WithHarmony)
	h.Any("/getidentity", m.WithRateLimit(h.GetIdentity, 2*time.Second, 20))
	a := h.Group.Group("", m.WithAuth)
	a.POST("/getguilds", m.WithRateLimit(h.GetGuilds, 5*time.Second, 5))
	a.POST("/getmembers", m.WithRateLimit(h.GetMembers, 5*time.Second, 3))
	a.POST("/getchannels", m.WithRateLimit(h.GetChannels, 500*time.Millisecond, 5))
	a.POST("/getmessages", m.WithRateLimit(h.GetMessages, 5*time.Second, 15))
	a.POST("/updateguildname", m.WithRateLimit(h.UpdateGuildName, 2*time.Second, 3))
	a.POST("/updateguildpicture", m.WithRateLimit(h.UpdateGuildPicture, 3*time.Second, 3))
	a.POST("/message", m.WithRateLimit(h.Message, 500*time.Millisecond, 20))
	a.POST("/createguild", m.WithRateLimit(h.CreateGuild, 20*time.Second, 3))
	a.POST("/addchannel", m.WithRateLimit(h.AddChannel, 1*time.Second, 3))
	a.POST("/deletechannel", m.WithRateLimit(h.DeleteChannel, 1*time.Second, 3))
	a.POST("/deleteguild", m.WithRateLimit(h.DeleteGuild, 5*time.Second, 5))
	a.POST("/deletemessage", m.WithRateLimit(h.DeleteMessage, 1*time.Second, 10))
	a.POST("/deleteinvite", m.WithRateLimit(h.DeleteInvite, 1*time.Second, 5))
	a.POST("/createinvite", m.WithRateLimit(h.CreateInvite, 1*time.Second, 5))
	a.POST("/getinvites", m.WithRateLimit(h.GetInvites, 2*time.Second, 4))
	a.POST("/joinguild", m.WithRateLimit(h.JoinGuild, 3*time.Second, 3))
	a.POST("/leaveguild", m.WithRateLimit(h.LeaveGuild, 3*time.Second, 2))

	return h
}
