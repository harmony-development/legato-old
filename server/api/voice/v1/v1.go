package v1

import (
	"encoding/json"
	"fmt"
	"time"

	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/api/voice/state"
	"github.com/harmony-development/legato/server/db"
	"github.com/pion/webrtc/v3"
	"github.com/ztrue/tracerr"
)

type Dependencies struct {
	DB         db.IHarmonyDB
	VoiceState *state.VoiceState
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
		Auth:     true,
		Location: middleware.GuildLocation | middleware.ChannelLocation,
	}, "/protocol.chat.v1.VoiceService/Connect")
}

func (v1 *V1) Connect(r *voicev1.ConnectRequest, s voicev1.VoiceService_ConnectServer) error {
	userID, err := middleware.AuthHandler(v1.DB, s.Context())
	if err := middleware.LocationHandler(v1.DB, r, "/protocol.chat.v1.VoiceService/Connect", userID); err != nil {
		fmt.Println(err)
		return err
	}
	if err != nil {
		return err
	}
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		return err
	}
	defer func() {
		if err := peerConnection.Close(); err != nil {
			fmt.Println("noo", err)
		}
	}()

	if _, err := peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio, webrtc.RTPTransceiverInit{
		Direction: webrtc.RTPTransceiverDirectionRecvonly,
	}); err != nil {
		return err
	}

	peerConnection.OnICECandidate(func(i *webrtc.ICECandidate) {
		if i == nil {
			return
		}
		candidateString, err := json.Marshal(i.ToJSON())
		if err != nil {
			fmt.Println(tracerr.Wrap(err))
			return
		}

		s.Send(&voicev1.Signal{
			Event: &voicev1.Signal_Candidate{
				Candidate: &voicev1.Signal_ICECandidate{
					Candidate: string(candidateString),
				},
			},
		})
	})

	return nil
}
