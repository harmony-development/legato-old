package state

import (
	"encoding/json"
	"sync"

	voicev1 "github.com/harmony-development/legato/gen/voice/v1"
	"github.com/pion/webrtc/v3"
)

type UserState struct {
	sync.RWMutex
	peerConnection *webrtc.PeerConnection
	stream         voicev1.VoiceService_ConnectServer
}

type Channel struct {
	sync.RWMutex
	Users  map[string]*UserState
	Tracks map[string]*webrtc.TrackLocalStaticRTP
}

type VoiceState struct {
	sync.RWMutex
	VoiceChannels map[string]*Channel
}

func (v *VoiceState) GetChannel(channelID string) *Channel {
	v.Lock()
	defer v.Unlock()
	if _, ok := v.VoiceChannels[channelID]; !ok {
		v.VoiceChannels[channelID] = &Channel{
			Users:   map[string]*UserState{},
			Tracks:  map[string]*webrtc.TrackLocalStaticRTP{},
			RWMutex: sync.RWMutex{},
		}
	}

	return v.VoiceChannels[channelID]
}

func (c *Channel) GetUserState(userID string) *UserState {
	c.Lock()
	defer c.Unlock()
	return c.Users[userID]
}

func (c *Channel) GetTracks() map[string]*webrtc.TrackLocalStaticRTP {
	c.Lock()
	defer c.Unlock()
	return c.Tracks
}

func (c *Channel) AddTrack(userID string, track *webrtc.TrackRemote) (*webrtc.TrackLocalStaticRTP, error) {
	c.Lock()
	defer c.Unlock()
	localTrack, err := webrtc.NewTrackLocalStaticRTP(track.Codec().RTPCodecCapability, track.ID(), track.StreamID())
	if err != nil {
		return nil, err
	}
	c.Tracks[userID] = localTrack
	for _, user := range c.Users {
		if _, err := user.peerConnection.AddTrack(localTrack); err != nil {
			return nil, err
		}
		offer, err := user.peerConnection.CreateOffer(nil)
		if err != nil {
			return nil, err
		}
		if err := user.peerConnection.SetLocalDescription(offer); err != nil {
			return nil, err
		}
		offerString, err := json.Marshal(offer)
		if err != nil {
			return nil, err
		}
		user.stream.Send(&voicev1.Signal{
			Event: &voicev1.Signal_Offer_{
				Offer: &voicev1.Signal_Offer{
					Offer: string(offerString),
				},
			},
		})
	}
	return localTrack, nil
}

func (c *Channel) RemoveTrack(userID string) {
	c.Lock()
	defer c.Unlock()
	delete(c.Tracks, userID)
}
