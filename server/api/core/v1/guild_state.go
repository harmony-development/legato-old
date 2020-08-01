package v1

import (
	"sync"

	corev1 "github.com/harmony-development/legato/gen/core"
)

type GuildState struct {
	serverChannels map[corev1.CoreService_StreamGuildEventsServer]chan struct{}
	guildEvents    map[UserID]map[GuildID][]corev1.CoreService_StreamGuildEventsServer
	subs           map[GuildID]map[UserID]struct{}
	sync.Mutex
}

func (s *GuildState) Add(guildID, userID uint64, server corev1.CoreService_StreamGuildEventsServer) chan struct{} {
	s.Lock()
	defer s.Unlock()

	s.SubAdd(guildID, userID)

	if _, ok := s.guildEvents[UserID(userID)]; !ok {
		s.guildEvents[UserID(userID)] = map[GuildID][]corev1.CoreService_StreamGuildEventsServer{}
	}

	go func() {
		<-server.Context().Done()
		s.RemoveUserFromGuild(userID, guildID)
	}()

	val, _ := s.guildEvents[UserID(userID)][GuildID(guildID)]
	val = append(val, server)
	s.guildEvents[UserID(userID)][GuildID(guildID)] = val
	s.serverChannels[server] = make(chan struct{})
	return s.serverChannels[server]
}

func (s *GuildState) Remove(userID uint64) {
	s.Lock()
	defer s.Unlock()

	for _, guild := range s.guildEvents[UserID(userID)] {
		for _, serv := range guild {
			close(s.serverChannels[serv])
			delete(s.serverChannels, serv)
		}
	}
	delete(s.guildEvents, UserID(userID))
	s.SubRemoveUser(userID)
}

func (s *GuildState) RemoveGuild(guildID uint64) {
	s.Lock()
	defer s.Unlock()
	defer delete(s.subs, GuildID(guildID))

	if val, ok := s.subs[GuildID(guildID)]; ok {
		for user := range val {
			for _, guild := range s.guildEvents[user] {
				for _, serv := range guild {
					close(s.serverChannels[serv])
					delete(s.serverChannels, serv)
				}
			}
			delete(s.guildEvents[user], GuildID(guildID))
		}
	}
}

func (s *GuildState) RemoveUserFromGuild(userID, guildID uint64) {
	s.Lock()
	defer s.Unlock()

	s.SubRemoveUserFromGuild(userID, guildID)
	if _, ok := s.guildEvents[UserID(userID)]; ok {
		for _, serv := range s.guildEvents[UserID(userID)][GuildID(guildID)] {
			close(s.serverChannels[serv])
			delete(s.serverChannels, serv)
		}
		delete(s.guildEvents[UserID(userID)], GuildID(guildID))
	}
}

func (s *GuildState) BroadcastGuild(guildID uint64, event *corev1.GuildEvent) {
	s.Lock()
	defer s.Unlock()

	go func() {
		for sub := range s.subs[GuildID(guildID)] {
			for _, server := range s.guildEvents[sub][GuildID(guildID)] {
				server.Send(event)
			}
		}
	}()
}
