package v1

import (
	"time"

	"github.com/sony/sonyflake"

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
	DB             *db.HarmonyDB
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	State          *state.State
	Sonyflake      *sonyflake.Sonyflake
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
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    5,
			},
			Auth:   true,
			Schema: nil,
		},
		{
			Path:    "/guilds",
			Handler: h.CreateGuild,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 20 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: new(CreateGuildData),
		},
		{
			Path:    "/guilds",
			Handler: h.DeleteGuild,
			Method:  routing.DELETE,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    5,
			},
			Auth:   true,
			Schema: new(DeleteGuildData),
		},
		{
			Path:    "/members",
			Handler: h.GetMembers,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: new(GetMembersData),
		},
		{
			Path:    "/channels",
			Handler: h.GetChannels,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    5,
			},
			Auth:   true,
			Schema: new(GetChannelsData),
		},
		{
			Path:    "/channels",
			Handler: h.AddChannel,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: new(AddChannelData),
		},
		{
			Path:    "/channels",
			Handler: h.DeleteChannel,
			Method:  routing.DELETE,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: new(DeleteChannelData),
		},
		{
			Path:    "/messages",
			Handler: h.GetMessages,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    15,
			},
			Auth:   true,
			Schema: new(GetMessagesData),
		},
		{
			Path:    "/guildname",
			Handler: h.UpdateGuildName,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 2 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: new(UpdateGuildNameData),
		},
		{
			Path:    "/guildpicture",
			Handler: h.UpdateGuildName,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: new(UpdateGuildPictureData),
		},
		{
			Path:    "/message",
			Handler: h.UpdateGuildName,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    20,
			},
			Auth:   true,
			Schema: new(UpdateGuildNameData),
		},
		{
			Path:    "/message",
			Handler: h.DeleteMessage,
			Method:  routing.DELETE,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    10,
			},
			Auth:   true,
			Schema: new(DeleteMessageData),
		},
		{
			Path:    "/invite",
			Handler: h.CreateInvite,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:   true,
			Schema: new(CreateInviteData),
		},
		{
			Path:    "/invite",
			Handler: h.DeleteInvite,
			Method:  routing.DELETE,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:   true,
			Schema: new(DeleteInviteData),
		},
		{
			Path:    "/invite",
			Handler: h.GetInvites,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 2 * time.Second,
				Burst:    4,
			},
			Auth:   true,
			Schema: new(DeleteInviteData),
		},
		{
			Path:    "/join",
			Handler: h.JoinGuild,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: new(JoinGuildData),
		},
		{
			Path:    "/leave",
			Handler: h.LeaveGuild,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    2,
			},
			Auth:   true,
			Schema: new(LeaveGuildData),
		},
		{
			Path:    "/key",
			Handler: h.GetKey,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    1024,
			},
			Auth:   false,
			Schema: nil,
		},
		{
			Path:    "/connect",
			Handler: h.Connect,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    1024,
			},
			Auth:   false,
			Schema: new(ConnectData),
		},
	}
}
