package webrtc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	m.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000))

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
	if _, err := peerConnection.AddTransceiver(webrtc.RTPCodecTypeAudio); err != nil {
		fmt.Println("error adding transceiver", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	peerConnection.OnTrack(api.OnTrackStart(peerConnection))

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

	return ctx.JSON(http.StatusOK, answer)
}

// OnTrackStart handles when a track is being received from a peer
func (api API) OnTrackStart(peerConnection *webrtc.PeerConnection) func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
	return func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		_, err := peerConnection.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), "audio", "useridhere")
		if err != nil {
			return
		}
		return
	}
}
