package voicebackend

import (
	"fmt"
	"sync"

	"github.com/pion/webrtc/v3"
)

type VoiceChannel struct {
	sync.RWMutex
	tracks          map[uint64]webrtc.TrackLocal
	peerConnections map[uint64]*Peer
}

type Peer struct {
	pc          *webrtc.PeerConnection
	addedTracks map[uint64]*webrtc.RTPSender
}

func (vc *VoiceChannel) NewPeer(userID uint64) (*webrtc.PeerConnection, error) {
	pc, err := webrtc.NewPeerConnection(peerConnectionConfig)
	if err != nil {
		return nil, err
	}
	vc.Lock()
	vc.peerConnections[userID] = &Peer{
		pc:          pc,
		addedTracks: map[uint64]*webrtc.RTPSender{},
	}
	vc.Unlock()
	if _, err := pc.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio, webrtc.RTPTransceiverInit{
		Direction: webrtc.RTPTransceiverDirectionSendrecv,
	}); err != nil {
		return nil, err
	}
	for _, track := range vc.tracks {
		pc.AddTrack(track)
	}
	return pc, nil
}

func (vc *VoiceChannel) DeletePeer(userID uint64) {
	vc.Lock()
	for _, peer := range vc.peerConnections {
		peer.pc.RemoveTrack(peer.addedTracks[userID])
	}
	delete(vc.peerConnections, userID)
	delete(vc.tracks, userID)
	vc.Unlock()
}

func (vc *VoiceChannel) AddTrack(userID uint64, track webrtc.TrackLocal) {
	vc.Lock()
	vc.tracks[userID] = track
	for _, conn := range vc.peerConnections {
		if rtpSender, err := conn.pc.AddTrack(track); err != nil {
			fmt.Println(err)
		} else {
			conn.addedTracks[userID] = rtpSender
		}
	}
	vc.Unlock()
}
