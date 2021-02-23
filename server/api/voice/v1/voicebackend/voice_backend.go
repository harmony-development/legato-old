package voicebackend

import (
	"sync"

	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/pion/webrtc/v3"
)

type VoiceBackend struct {
	activeChannels map[uint64]*VoiceChannel
	users          map[uint64]*chan *voicev1.Signal
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

func (backend *VoiceBackend) RegisterUserStream(userID uint64, stream chan *voicev1.Signal) {
	backend.users[userID] = &stream
}

func (backend *VoiceBackend) GetStream(userID uint64) *chan *voicev1.Signal {
	return backend.users[userID]
}
