package v1

import (
	"fmt"
	"time"

	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/api/voice/state"
	"github.com/harmony-development/legato/server/db"
	"github.com/pion/webrtc/v3"
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
		Location: middleware.GuildLocation,
	}, "/protocol.chat.v1.VoiceService/Connect")
}

func (v1 *V1) Connect(s voicev1.VoiceService_ConnectServer) error {
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

	return nil
}
