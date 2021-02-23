package v1

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/api/voice/v1/voicebackend"
	"github.com/harmony-development/legato/server/db/types"
	"github.com/labstack/echo/v4"
	"github.com/pion/rtcp"
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

func (v1 *V1) onTrack(userID, channelID uint64, peer *webrtc.PeerConnection) func(*webrtc.TrackRemote, *webrtc.RTPReceiver) {
	return func(remoteTrack *webrtc.TrackRemote, r *webrtc.RTPReceiver) {
		vc := v1.VoiceBackend.GetVoiceChannel(channelID)
		if vc == nil {
			return
		}

		// send a PLI to force keyframes to be pushed
		go func() {
			ticker := time.NewTicker(time.Second * 3)
			for range ticker.C {
				if rtcpSendErr := peer.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: uint32(remoteTrack.SSRC())}}); rtcpSendErr != nil {
					fmt.Println(rtcpSendErr)
				}
			}
		}()

		// make a local track that feeds into all clients
		localTrack, err := webrtc.NewTrackLocalStaticRTP(remoteTrack.Codec().RTPCodecCapability, remoteTrack.ID(), remoteTrack.StreamID())
		if err != nil {
			return
		}

		vc.AddTrack(userID, localTrack)
		for _, stateStream := range vc.StateStreams {
			stateStream <- &voicev1.Signal{
				Event: &voicev1.Signal_RenegotiationNeeded{
					RenegotiationNeeded: &empty.Empty{},
				},
			}
		}

		buf := make([]byte, 1500)
		for {
			i, _, err := remoteTrack.Read(buf)
			if err != nil {
				vc.DeletePeer(userID)
			}
			if _, err = localTrack.Write(buf[:i]); err != nil {
				return
			}
		}
	}
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

	peer.OnTrack(v1.onTrack(ctx.UserID, r.ChannelId, peer))

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
}
