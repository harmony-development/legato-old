package webrtc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/labstack/echo/v4"
	"github.com/pion/webrtc/v3"
)

// SDPHandler handles webrtc peer connection attempts
func (api API) SDPHandler(c echo.Context) error {
	ctx := c.(hm.HarmonyContext)
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.NoContent(http.StatusUnprocessableEntity)
	}

	m := webrtc.MediaEngine{}
	m.RegisterCodec(webrtc.NewRTPOpusCodec(webrtc.DefaultPayloadTypeOpus, 48000))

	mediaAPI := webrtc.NewAPI(webrtc.WithMediaEngine(m))

	offer := webrtc.SessionDescription{}
	if err := json.Unmarshal(body, &offer); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	peerConnection, err := mediaAPI.NewPeerConnection(api.peerConnectionConfig)
	if err != nil {
		fmt.Println("error making peer connection", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	if _, err := peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio); err != nil {
		fmt.Println("error adding transceiver", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	peerConnection.OnTrack(api.OnTrackStart(peerConnection, *ctx.Location.ChannelID, ctx.UserID))
	peerConnection.OnICEConnectionStateChange(api.OnICEConnectionStateChange(peerConnection, *ctx.Location.ChannelID, ctx.UserID))

	if err := peerConnection.SetRemoteDescription(offer); err != nil {
		fmt.Println("error setting remote description", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		fmt.Println("error making answer", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	if err := peerConnection.SetLocalDescription(answer); err != nil {
		fmt.Println("error setting local description", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	if _, exists := api.VoiceTracks[*ctx.Location.ChannelID]; !exists {
		api.VoiceTracks[*ctx.Location.ChannelID] = make(map[uint64]*webrtc.Track)
	}

	for userID := range api.VoiceTracks[*ctx.Location.ChannelID] {
		peerConnection.AddTrack(api.VoiceTracks[*ctx.Location.ChannelID][userID])
	}

	return ctx.JSON(http.StatusOK, answer)
}

// OnTrackStart handles when a track is being received from a peer
func (api API) OnTrackStart(peerConnection *webrtc.PeerConnection, channelID, userID uint64) func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
	return func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		track, err := peerConnection.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), strconv.FormatUint(userID, 10), strconv.FormatUint(userID, 10))
		if err != nil {
			return
		}
		api.VoiceTracks[channelID][userID] = track
		fmt.Println(api.VoiceTracks)
		return
	}
}

// OnICEConnectionStateChange handles webrtc state changes such as timeouts
func (api API) OnICEConnectionStateChange(peerConnection *webrtc.PeerConnection, channelID, userID uint64) func(webrtc.ICEConnectionState) {
	return func(state webrtc.ICEConnectionState) {
		if state == webrtc.ICEConnectionStateDisconnected || state == webrtc.ICEConnectionStateClosed {
			if err := peerConnection.Close(); err != nil {
				fmt.Println(err)
			}
			delete(api.VoiceTracks[channelID], userID)
			if len(api.VoiceTracks[channelID]) == 0 {
				delete(api.VoiceTracks, channelID)
			}
		}
	}
}
