package webrtc

import (
	"encoding/json"
	"fmt"
	"io"
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
	channelID := *ctx.Location.ChannelID
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.NoContent(http.StatusUnprocessableEntity)
	}

	if api.VoiceChannels[channelID] == nil {
		api.VoiceChannels[channelID] = &VoiceChannel{
			Tracks: make(map[uint64]*webrtc.Track),
			Peers:  make(map[uint64]*webrtc.PeerConnection),
		}
	}

	offer := webrtc.SessionDescription{}
	if err := json.Unmarshal(body, &offer); err != nil {
		fmt.Println("error parsing SDP", err)
		return ctx.NoContent(http.StatusBadRequest)
	}

	peerConnection, err := api.MediaAPI.NewPeerConnection(api.peerConnectionConfig)
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

	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	if err := peerConnection.SetLocalDescription(answer); err != nil {
		fmt.Println("error setting local description", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	<-gatherComplete

	if _, exists := api.VoiceChannels[*ctx.Location.ChannelID]; !exists {
		api.VoiceChannels[*ctx.Location.ChannelID] = &VoiceChannel{
			Tracks: make(map[uint64]*webrtc.Track),
			Peers:  make(map[uint64]*webrtc.PeerConnection),
		}
	}

	for userID := range api.VoiceChannels[*ctx.Location.ChannelID].Tracks {
		if _, err := peerConnection.AddTrack(api.VoiceChannels[*ctx.Location.ChannelID].Tracks[userID]); err != nil {
			fmt.Println(err)
		}
	}

	api.VoiceChannels[*ctx.Location.ChannelID].Peers[ctx.UserID] = peerConnection

	return ctx.JSON(http.StatusOK, peerConnection.LocalDescription())
}

// OnTrackStart handles when a track is being received from a peer
func (api API) OnTrackStart(peerConnection *webrtc.PeerConnection, channelID, userID uint64) func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
	return func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		localTrack, err := peerConnection.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), strconv.FormatUint(userID, 10), strconv.FormatUint(userID, 10))
		if err != nil {
			fmt.Println(err)
			return
		}
		api.VoiceChannels[channelID].Tracks[userID] = localTrack
		for userID := range api.VoiceChannels[channelID].Peers {
			if _, err := api.VoiceChannels[channelID].Peers[userID].AddTrack(localTrack); err != nil {
				fmt.Println(err)
			}
		}
		rtpBuf := make([]byte, 512)
		for {
			i, readErr := remoteTrack.Read(rtpBuf)
			if readErr != nil {
				fmt.Println(readErr)
				return
			}
			if _, err = localTrack.Write(rtpBuf[:i]); err != nil && err != io.ErrClosedPipe {
				fmt.Println("oops")
				return
			}
		}
	}
}

// OnICEConnectionStateChange handles webrtc state changes such as timeouts
func (api API) OnICEConnectionStateChange(peerConnection *webrtc.PeerConnection, channelID, userID uint64) func(webrtc.ICEConnectionState) {
	return func(state webrtc.ICEConnectionState) {
		if state == webrtc.ICEConnectionStateDisconnected || state == webrtc.ICEConnectionStateClosed {
			fmt.Println("disconnect", channelID, userID)
			if err := peerConnection.Close(); err != nil {
				fmt.Println(err)
			}
			delete(api.VoiceChannels[channelID].Tracks, userID)
			if len(api.VoiceChannels[channelID].Tracks) == 0 {
				delete(api.VoiceChannels, channelID)
			}
		}
	}
}
