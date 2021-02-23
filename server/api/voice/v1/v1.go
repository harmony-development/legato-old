package v1

import (
	"encoding/json"
	"time"

	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/api/voice/voicebackend"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/labstack/echo/v4"
	"github.com/pion/webrtc/v3"
)

type Dependencies struct {
	DB           types.IHarmonyDB
	VoiceBackend *voicebackend.VoiceBackend
}

type V1 struct {
	Dependencies
}

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    5,
		},
		Location: middleware.ChannelLocation,
	}, "/protocol.chat.v1.VoiceService/Connect")
}

func (v1 *V1) Connect(c echo.Context, r *voicev1.ConnectRequest) (*voicev1.ConnectResponse, error) {
	ctx := c.(middleware.HarmonyContext)
	vc := v1.VoiceBackend.GetVoiceChannel(r.ChannelId)

	offer := webrtc.SessionDescription{}
	if err := json.Unmarshal([]byte(r.Offer), &offer); err != nil {
		return nil, err
	}

	peer, err := vc.NewPeer(ctx.UserID)
	if err != nil {
		return nil, err
	}

	if err := peer.SetRemoteDescription(offer); err != nil {
		return nil, err
	}

	answer, err := peer.CreateAnswer(nil)
	if err != nil {
		return nil, err
	}

	// TODO! change this to use OnICECandidate before going into production
	gatherComplete := webrtc.GatheringCompletePromise(peer)

	err = peer.SetLocalDescription(answer)
	if err != nil {
		return nil, err
	}

	<-gatherComplete

	marshalled, err := json.Marshal(peer.LocalDescription())
	if err != nil {
		return nil, err
	}

	return &voicev1.ConnectResponse{
		Answer: string(marshalled),
	}, nil
}

func (v1 *V1) StreamState(c echo.Context, r *voicev1.StreamStateRequest, out chan *voicev1.Signal) {
	return
}
