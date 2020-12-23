package integrated

import (
	"sync"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
)

// GuildState is the state of a guild
type GuildState struct {
	serverChannels map[chatv1.ChatService_StreamEventsServer]chan struct{}
	guildEvents    map[_userID]map[_guildID][]chatv1.ChatService_StreamEventsServer
	subs           map[_guildID]map[_userID]struct{}
	sync.Mutex
}

// Initialize the guild state
func (s *GuildState) Initialize() *GuildState {
	s.serverChannels = make(map[chatv1.ChatService_StreamEventsServer]chan struct{})
	s.guildEvents = make(map[_userID]map[_guildID][]chatv1.ChatService_StreamEventsServer)
	s.subs = make(map[_guildID]map[_userID]struct{})
	return s
}

// Subscribe ...
func (s *GuildState) Subscribe(guildID, userID uint64, server chatv1.ChatService_StreamEventsServer) chan struct{} {
	s.Lock()
	defer s.Unlock()

	s.subAdd(guildID, userID)

	if _, ok := s.guildEvents[_userID(userID)]; !ok {
		s.guildEvents[_userID(userID)] = map[_guildID][]chatv1.ChatService_StreamEventsServer{}
	}

	go func() {
		<-server.Context().Done()
		s.UnsubscribeUserFromGuild(userID, guildID)
	}()
	val, ok := s.guildEvents[_userID(userID)][_guildID(guildID)]
	_ = ok
	val = append(val, server)
	s.guildEvents[_userID(userID)][_guildID(guildID)] = val
	s.serverChannels[server] = make(chan struct{})
	return s.serverChannels[server]
}

// UnsubscribeUser ...
func (s *GuildState) UnsubscribeUser(userID uint64) {
	s.Lock()
	defer s.Unlock()

	for _, guild := range s.guildEvents[_userID(userID)] {
		for _, serv := range guild {
			close(s.serverChannels[serv])
			delete(s.serverChannels, serv)
		}
	}
	delete(s.guildEvents, _userID(userID))
	s.subRemoveUser(userID)
}

// UnsubscribeGuild ...
func (s *GuildState) UnsubscribeGuild(guildID uint64) {
	s.Lock()
	defer s.Unlock()
	defer delete(s.subs, _guildID(guildID))

	if val, ok := s.subs[_guildID(guildID)]; ok {
		for user := range val {
			for _, guild := range s.guildEvents[user] {
				for _, serv := range guild {
					if _, ok := s.serverChannels[serv]; ok {
						close(s.serverChannels[serv])
						delete(s.serverChannels, serv)
					}
				}
			}
			delete(s.guildEvents[user], _guildID(guildID))
		}
	}
}

// UnsubscribeUserFromGuild ...
func (s *GuildState) UnsubscribeUserFromGuild(userID, guildID uint64) {
	s.Lock()
	defer s.Unlock()

	s.subRemoveUserFromGuild(userID, guildID)
	if _, ok := s.guildEvents[_userID(userID)]; ok {
		for _, serv := range s.guildEvents[_userID(userID)][_guildID(guildID)] {
			if _, ok := s.serverChannels[serv]; ok {
				close(s.serverChannels[serv])
				delete(s.serverChannels, serv)
			}
		}
		delete(s.guildEvents[_userID(userID)], _guildID(guildID))
	}
}

// Broadcast ...
func (s *GuildState) Broadcast(guildID uint64, event *chatv1.Event) {
	s.Lock()
	defer s.Unlock()

	go func() {
		for sub := range s.subs[_guildID(guildID)] {
			for _, server := range s.guildEvents[sub][_guildID(guildID)] {
				if err := server.Send(event); err != nil {
					println(err)
				}
			}
		}
	}()
}
