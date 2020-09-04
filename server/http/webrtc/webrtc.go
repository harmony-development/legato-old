package webrtc

import (
	"time"

	"github.com/harmony-development/legato/server/http/routing"
	"github.com/labstack/echo/v4"
	"github.com/pion/webrtc/v3"
)

type Dependencies struct {
	APIGroup *echo.Group
	Router   routing.IRouter
}

type API struct {
	*echo.Group
	Dependencies
	peerConnectionConfig webrtc.Configuration
}

func New(deps Dependencies) *API {
	api := &API{
		Group:        deps.APIGroup.Group("/webrtc"),
		Dependencies: deps,
		peerConnectionConfig: webrtc.Configuration{
			ICEServers: []webrtc.ICEServer{
				{
					URLs: []string{"stun:stun.l.google.com:19302"},
				},
			},
		},
	}

	api.Router.BindRoutes(api.Group, []routing.Route{
		{
			Path:    "/:guild_id/:channel_id/sdp",
			Handler: api.SDPHandler,
			Auth:    true,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    6,
			},
			Location: routing.LocationGuildAndChannel,
			Method:   routing.POST,
		},
	})
	return api
}
