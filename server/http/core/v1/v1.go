package v1

import (
	"time"

	"github.com/sony/sonyflake"

	"harmony-server/server/auth"
	"harmony-server/server/config"
	"harmony-server/server/db"
	"harmony-server/server/http/hm"
	"harmony-server/server/http/routing"
	"harmony-server/server/logger"
	"harmony-server/server/state"
	"harmony-server/server/storage"
)

// Handlers for CoreKit
type Handlers struct {
	Deps *Dependencies
}

// Dependencies are the elements that CoreKit handlers need
type Dependencies struct {
	DB             db.IHarmonyDB
	Config         *config.Config
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Logger         *logger.Logger
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
			Handler: h.CreateGuild,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 20 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: CreateGuildData{},
		},
		{
			Path:    "/guilds/:guild_id/name",
			Handler: h.UpdateGuildName,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 2 * time.Second,
				Burst:    3,
			},
			Auth:        true,
			Schema:      UpdateGuildNameData{},
			Location:    routing.LocationGuild,
			Permissions: hm.ModifyGuild,
		},
		{
			Path:    "/guilds/:guild_id/picture",
			Handler: h.UpdateGuildPicture,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 2 * time.Second,
				Burst:    3,
			},
			Auth:        true,
			Location:    routing.LocationGuild,
			Permissions: hm.ModifyGuild,
		},
		{
			Path:    "/guilds/:guild_id",
			Handler: h.DeleteGuild,
			Method:  routing.DELETE,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    5,
			},
			Auth:        true,
			Location:    routing.LocationGuild,
			Permissions: hm.Owner,
		},
		{
			Path:    "/guilds/:guild_id",
			Handler: h.GetGuild,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    5,
			},
			Auth:     true,
			Location: routing.LocationGuild,
		},
		{
			Path:    "/guilds/:guild_id/members",
			Handler: h.GetMembers,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    3,
			},
			Auth:     true,
			Location: routing.LocationGuild,
		},
		{
			Path:    "/guilds/:guild_id/channels",
			Handler: h.GetChannels,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    5,
			},
			Auth:     true,
			Location: routing.LocationGuild,
		},
		{
			Path:    "/guilds/:guild_id/channels",
			Handler: h.AddChannel,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Auth:        true,
			Schema:      AddChannelData{},
			Location:    routing.LocationGuild,
			Permissions: hm.ModifyChannels,
		},
		{
			Path:    "/guilds/:guild_id/channels/:channel_id",
			Handler: h.DeleteChannel,
			Method:  routing.DELETE,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Auth:        true,
			Location:    routing.LocationGuildAndChannel,
			Permissions: hm.ModifyChannels,
		},
		{
			Path:    "/guilds/:guild_id/invites",
			Handler: h.CreateInvite,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:        true,
			Schema:      CreateInviteData{},
			Location:    routing.LocationGuild,
			Permissions: hm.ModifyInvites,
		},
		{
			Path:    "/guilds/:guild_id/invites",
			Handler: h.GetInvites,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 2 * time.Second,
				Burst:    4,
			},
			Auth:        true,
			Location:    routing.LocationGuild,
			Permissions: hm.ModifyInvites,
		},
		{
			Path:    "/guilds/:guild_id/invites/:invite_id",
			Handler: h.DeleteInvite,
			Method:  routing.DELETE,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:        true,
			Location:    routing.LocationGuild,
			Permissions: hm.ModifyInvites,
		},
		{
			Path:    "/guilds/:guild_id/channels/:channel_id/messages",
			Handler: h.GetMessages,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    15,
			},
			Auth:     true,
			Schema:   GetMessagesData{},
			Location: routing.LocationGuildAndChannel,
		},
		{
			Path:    "/guilds/:guild_id/channels/:channel_id/messages",
			Handler: h.Message,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    20,
			},
			Auth:     true,
			Schema:   MessageData{},
			Location: routing.LocationGuildAndChannel,
		},
		{
			Path:    "/guilds/:guild_id/channels/:channel_id/messages/:message_id",
			Handler: h.UpdateMessage,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    20,
			},
			Auth:     true,
			Schema:   MessageUpdateData{},
			Location: routing.LocationGuildChannelAndMessage,
		},
		{
			Path:    "/guilds/:guild_id/channels/:channel_id/messages/:message_id",
			Handler: h.DeleteMessage,
			RateLimit: &routing.RateLimit{
				Duration: 500 * time.Millisecond,
				Burst:    20,
			},
			Auth:     true,
			Location: routing.LocationGuildChannelAndMessage,
		},
		{
			Path:    "/guilds/:guild_id/channels/:channel_id/messages/:message_id/trigger_action",
			Handler: h.TriggerAction,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:     true,
			Schema:   TriggerActionData{},
			Location: routing.LocationGuildChannelAndMessage,
		},
		{
			Path:    "/users/~/guilds",
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
			Path:    "/users/~/guilds/join",
			Handler: h.JoinGuild,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    3,
			},
			Auth:   true,
			Schema: JoinGuildData{},
		},
		{
			Path:    "/users/~/guilds/leave/:guild_id",
			Handler: h.LeaveGuild,
			Method:  routing.POST,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    2,
			},
			Auth:     true,
			Schema:   LeaveGuildData{},
			Location: routing.LocationGuild,
		},
	}
}
