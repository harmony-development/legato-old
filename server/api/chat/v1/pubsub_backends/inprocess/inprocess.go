package inprocess

import (
	"sync"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/logger"
)

func unused(v interface{}) {

}

// streamData is the data of a stream
type streamData struct {
	userID uint64

	guilds     map[uint64]struct{}
	homeserver bool
	action     bool
}

// StreamManager manages streams
type StreamManager struct {
	logger   logger.ILogger
	database db.IHarmonyDB

	serverToStreamData map[chatv1.ChatService_StreamEventsServer]streamData
	userIDToServers    map[uint64]map[chatv1.ChatService_StreamEventsServer]struct{}
	guildIDToUserIDs   map[uint64]map[uint64]struct{}

	sync.RWMutex
}

// Init prepares a stream manager for use
func (s *StreamManager) Init(l logger.ILogger, db db.IHarmonyDB) {
	s.logger = l
	s.database = db

	s.serverToStreamData = make(map[chatv1.ChatService_StreamEventsServer]streamData)
	s.userIDToServers = make(map[uint64]map[chatv1.ChatService_StreamEventsServer]struct{})
	s.guildIDToUserIDs = make(map[uint64]map[uint64]struct{})
}

// RegisterClient registers a client
func (s *StreamManager) RegisterClient(userID uint64, srv chatv1.ChatService_StreamEventsServer) (ret chan struct{}) {
	s.Lock()
	defer s.Unlock()

	ret = make(chan struct{})

	s.serverToStreamData[srv] = streamData{
		userID: userID,
		guilds: make(map[uint64]struct{}),
	}

	servs, ok := s.userIDToServers[userID]
	unused(ok)
	if servs == nil {
		s.userIDToServers[userID] = make(map[chatv1.ChatService_StreamEventsServer]struct{})
		servs = s.userIDToServers[userID]
	}
	servs[srv] = struct{}{}

	s.userIDToServers[userID] = servs

	go func() {
		<-srv.Context().Done()

		s.Lock()
		defer s.Unlock()

		delete(s.serverToStreamData, srv)
		delete(s.userIDToServers[userID], srv)

		if len(s.userIDToServers[userID]) == 0 {
			delete(s.userIDToServers, userID)
		}

		close(ret)
	}()

	return ret
}

// AddGuildSubscription adds a subscription
func (s *StreamManager) AddGuildSubscription(srv chatv1.ChatService_StreamEventsServer, to uint64) {
	s.Lock()
	defer s.Unlock()

	s.serverToStreamData[srv].guilds[to] = struct{}{}

	g, ok := s.guildIDToUserIDs[to]
	unused(ok)
	if g == nil {
		g = make(map[uint64]struct{})
		s.guildIDToUserIDs[to] = g
	}

	g[to] = struct{}{}
}

// AddHomeserverSubscription adds a subscription
func (s *StreamManager) AddHomeserverSubscription(srv chatv1.ChatService_StreamEventsServer) {
	s.Lock()
	defer s.Unlock()

	strct := s.serverToStreamData[srv]
	strct.homeserver = true
	s.serverToStreamData[srv] = strct
}

// AddActionSubscription adds a subscription
func (s *StreamManager) AddActionSubscription(srv chatv1.ChatService_StreamEventsServer) {
	s.Lock()
	defer s.Unlock()

	strct := s.serverToStreamData[srv]
	strct.action = true
	s.serverToStreamData[srv] = strct
}

// RemoveGuildSubscription does a thing
func (s *StreamManager) RemoveGuildSubscription(userID, guildID uint64) {
	s.Lock()
	defer s.Unlock()

	delete(s.guildIDToUserIDs[guildID], userID)
	for serv := range s.userIDToServers[userID] {
		delete(s.serverToStreamData[serv].guilds, guildID)
	}
}

// BroadcastGuild does a thing
func (s *StreamManager) BroadcastGuild(to uint64, event *chatv1.Event) {
	go func() {
		s.RLock()
		defer s.RUnlock()

		for userID := range s.guildIDToUserIDs[to] {
			for serv := range s.userIDToServers[userID] {
				if _, ok := s.serverToStreamData[serv].guilds[to]; ok {
					err := serv.Send(event)
					unused(err)
				}
			}
		}
	}()
}

// BroadcastHomeserver does a thing
func (s *StreamManager) BroadcastHomeserver(userid uint64, event *chatv1.Event) {
	go func() {
		s.RLock()
		defer s.RUnlock()

		for server := range s.userIDToServers[userid] {
			if s.serverToStreamData[server].homeserver {
				err := server.Send(event)
				unused(err)
			}
		}
	}()
}

// BroadcastAction does a thing
func (s *StreamManager) BroadcastAction(userid uint64, event *chatv1.Event) {
	go func() {
		s.RLock()
		defer s.RUnlock()

		for server := range s.userIDToServers[userid] {
			if s.serverToStreamData[server].action {
				err := server.Send(event)
				unused(err)
			}
		}
	}()
}
