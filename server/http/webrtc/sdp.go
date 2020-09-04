package webrtc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/harmony-development/legato/server/http/hm"
	"github.com/labstack/echo/v4"
	"github.com/pion/webrtc/v3"
)

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
		return ctx.NoContent(http.StatusInternalServerError)
	}
	if _, err := peerConnection.AddTransceiver(webrtc.RTPCodecTypeAudio); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	peerConnection.OnTrack(api.OnTrackStart)

	if err := peerConnection.SetRemoteDescription(offer); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	if err := peerConnection.SetLocalDescription(answer); err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.JSON(http.StatusOK, answer)
}

func (api API) OnTrackStart(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {

}
