package v1

import (
	"time"

	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	"harmony-server/server/http/routing"
	"harmony-server/server/state"
	"harmony-server/server/storage"
)

// Handlers for CoreKit
type Handlers struct {
	Deps *Dependencies
}

// Dependencies are the elements that CoreKit handlers need
type Dependencies struct {
	DB             *db.DB
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	State          *state.State
}

// New creates a new set of Handlers
func New(deps *Dependencies) *Handlers {
	return &Handlers{
		Deps: deps,
	}
}

// MakeRoutes creates the routes for CoreKit
func (h Handlers) MakeRoutes() []routing.Route {
	return []routing.Route{
		{
			Path:    "/guilds",
			Handler: h.GetGuilds,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    5,
			},
			Auth:   true,
			Method: routing.GET,
		},
		{
			Path:    "/guilds",
			Handler: h.CreateGuild,
			RateLimit: &routing.RateLimit{
				Duration: 20 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Method: routing.POST,
		},
		{
			Path:    "/guilds",
			Handler: h.DeleteGuild,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    5,
			},
			Auth:   true,
			Method: routing.DELETE,
		},
		{
			Path:    "/members",
			Handler: h.GetMembers,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Method: routing.GET,
		},
		{
			Path:    "/channels",
			Handler: h.GetChannels,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    5,
			},
			Auth:   true,
			Method: routing.GET,
		},
		{
			Path:    "/channels",
			Handler: h.AddChannel,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Method: routing.POST,
		},
		{
			Path:    "/channels",
			Handler: h.DeleteChannel,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Method: routing.DELETE,
		},
		{
			Path:    "/messages",
			Handler: h.GetMessages,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    15,
			},
			Auth:   true,
			Method: routing.GET,
		},
		{
			Path:    "/guildname",
			Handler: h.UpdateGuildName,
			RateLimit: &routing.RateLimit{
				Duration: 2 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Method: routing.PATCH,
		},
		{
			Path:    "/guildpicture",
			Handler: h.UpdateGuildName,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Method: routing.PATCH,
		},
		{
			Path:    "/message",
			Handler: h.UpdateGuildName,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    20,
			},
			Auth:   true,
			Method: routing.POST,
		},
		{
			Path:    "/message",
			Handler: h.DeleteMessage,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    10,
			},
			Auth:   true,
			Method: routing.DELETE,
		},
		{
			Path:    "/invite",
			Handler: h.CreateInvite,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:   true,
			Method: routing.POST,
		},
		{
			Path:    "/invite",
			Handler: h.DeleteInvite,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:   true,
			Method: routing.DELETE,
		},
		{
			Path:    "/invite",
			Handler: h.GetInvites,
			RateLimit: &routing.RateLimit{
				Duration: 2 * time.Second,
				Burst:    4,
			},
			Auth:   true,
			Method: routing.GET,
		},
		{
			Path:    "/join",
			Handler: h.JoinGuild,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Method: routing.POST,
		},
		{
			Path:    "/leave",
			Handler: h.LeaveGuild,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    2,
			},
			Auth:   true,
			Method: routing.POST,
		},
		{
			Path:    "/key",
			Handler: h.GetKey,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    1024,
			},
			Auth:   false,
			Method: routing.GET,
		},
	}
}
