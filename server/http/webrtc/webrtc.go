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

type VoiceChannel struct {
	Tracks map[uint64]*webrtc.Track
	Peers  map[uint64]*webrtc.PeerConnection
}

type API struct {
	*echo.Group
	Dependencies
	peerConnectionConfig webrtc.Configuration
	Engine               webrtc.MediaEngine
	MediaAPI             *webrtc.API
	VoiceChannels        map[uint64]*VoiceChannel
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
		VoiceChannels: make(map[uint64]*VoiceChannel),
		Engine:        webrtc.MediaEngine{},
	}

	api.Engine.RegisterCodec(webrtc.NewRTPOpusCodec(webrtc.DefaultPayloadTypeOpus, 48000))

	api.MediaAPI = webrtc.NewAPI(webrtc.WithMediaEngine(api.Engine))

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
