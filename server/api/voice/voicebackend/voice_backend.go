package voicebackend

import (
	"sync"

	"github.com/pion/webrtc/v3"
)

type VoiceBackend struct {
	activeChannels map[uint64]*VoiceChannel
}

// TODO: make STUN server customizable
var peerConnectionConfig = webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs: []string{"stun:stun.l.google.com:19302"},
		},
	},
}

func (backend *VoiceBackend) GetVoiceChannel(channelID uint64) *VoiceChannel {
	if _, ok := backend.activeChannels[channelID]; !ok {
		backend.activeChannels[channelID] = &VoiceChannel{
			RWMutex:         sync.RWMutex{},
			tracks:          map[uint64]webrtc.TrackLocal{},
			peerConnections: map[uint64]*Peer{},
		}
	}
	return backend.activeChannels[channelID]
}
