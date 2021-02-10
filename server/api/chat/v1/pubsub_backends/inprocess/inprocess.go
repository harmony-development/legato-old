package inprocess

import (
	"sync"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/harmony-development/legato/server/db/types"
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
	database types.IHarmonyDB

	serverToStreamData map[chan *chatv1.Event]streamData
	userIDToServers    map[uint64]map[chan *chatv1.Event]struct{}
	guildIDToUserIDs   map[uint64]map[uint64]struct{}

	sync.RWMutex
}

// Init prepares a stream manager for use
func (s *StreamManager) Init(l logger.ILogger, db types.IHarmonyDB) {
	s.logger.Verbose(logger.Streams, "Initialising stream manager...")
	defer s.logger.Verbose(logger.Streams, "Initialising stream manager...")

	s.logger = l
	s.database = db

	s.serverToStreamData = make(map[chan *chatv1.Event]streamData)
	s.userIDToServers = make(map[uint64]map[chan *chatv1.Event]struct{})
	s.guildIDToUserIDs = make(map[uint64]map[uint64]struct{})
}

// RegisterClient registers a client
func (s *StreamManager) RegisterClient(userID uint64, srv chan *chatv1.Event, done chan struct{}) {
	s.Lock()
	defer s.Unlock()

	s.logger.Verbose(logger.Streams, "Registering user(%d) with server %+v", userID, srv)

	s.serverToStreamData[srv] = streamData{
		userID: userID,
		guilds: make(map[uint64]struct{}),
	}

	servs, ok := s.userIDToServers[userID]
	unused(ok)
	if servs == nil {
		s.logger.Verbose(logger.Streams, "Servers for user(%d) is nil; initialising...", userID)
		s.userIDToServers[userID] = make(map[chan *chatv1.Event]struct{})
		servs = s.userIDToServers[userID]
	}
	servs[srv] = struct{}{}

	s.userIDToServers[userID] = servs
	s.logger.Verbose(logger.Streams, "Registered server %+v for user(%d)", srv, userID)

	go func() {
		s.logger.Verbose(logger.Streams, "Waiting on user (%d)'s stream %+v to complete", userID, srv)
		<-done
		s.logger.Verbose(logger.Streams, "User (%d)'s stream %+v completed", userID, srv)

		s.Lock()
		defer s.Unlock()

		delete(s.serverToStreamData, srv)
		delete(s.userIDToServers[userID], srv)

		if len(s.userIDToServers[userID]) == 0 {
			s.logger.Verbose(logger.Streams, "User (%d) has no more open streams, deleting s.userIDToServers[userID]", userID, srv)
			delete(s.userIDToServers, userID)
		}
	}()
}

// AddGuildSubscription adds a subscription
func (s *StreamManager) AddGuildSubscription(srv chan *chatv1.Event, toGuild uint64) {
	s.Lock()
	defer s.Unlock()

	s.serverToStreamData[srv].guilds[toGuild] = struct{}{}

	userID := s.serverToStreamData[srv].userID

	s.logger.Verbose(logger.Streams, "Subscribing user(%d) server %+v to guild(%d)", userID, srv, toGuild)

	g, ok := s.guildIDToUserIDs[toGuild]
	unused(ok)
	if g == nil {
		g = make(map[uint64]struct{})
		s.guildIDToUserIDs[toGuild] = g
	}

	g[userID] = struct{}{}
}

// AddHomeserverSubscription adds a subscription
func (s *StreamManager) AddHomeserverSubscription(srv chan *chatv1.Event) {
	s.Lock()
	defer s.Unlock()

	s.logger.Verbose(logger.Streams, "Subscribing user(%d) server %+v to homeserver events", s.serverToStreamData[srv].userID, srv)

	strct := s.serverToStreamData[srv]
	strct.homeserver = true
	s.serverToStreamData[srv] = strct
}

// AddActionSubscription adds a subscription
func (s *StreamManager) AddActionSubscription(srv chan *chatv1.Event) {
	s.Lock()
	defer s.Unlock()

	s.logger.Verbose(logger.Streams, "Subscribing user(%d) server %+v to action events", s.serverToStreamData[srv].userID, srv)

	strct := s.serverToStreamData[srv]
	strct.action = true
	s.serverToStreamData[srv] = strct
}

// RemoveGuildSubscription does a thing
func (s *StreamManager) RemoveGuildSubscription(userID, guildID uint64) {
	s.Lock()
	defer s.Unlock()

	s.logger.Verbose(logger.Streams, "Unsubscribing user(%d) from guild %d", userID, guildID)

	delete(s.guildIDToUserIDs[guildID], userID)
	for serv := range s.userIDToServers[userID] {
		s.logger.Verbose(logger.Streams, "Deleting server %+v for %d", serv, userID)
		delete(s.serverToStreamData[serv].guilds, guildID)
	}
}

// BroadcastGuild does a thing
func (s *StreamManager) BroadcastGuild(to uint64, event *chatv1.Event) {
	s.logger.Verbose(logger.Streams, "Broadcasting guild event %+v to guild %d", event, to)

	go func() {
		s.RLock()
		defer s.RUnlock()

		for userID := range s.guildIDToUserIDs[to] {
			s.logger.Verbose(logger.Streams, "Broadcasting guild event %+v to user %d for guild %d", event, userID, to)
			for serv := range s.userIDToServers[userID] {
				if _, ok := s.serverToStreamData[serv].guilds[to]; ok {
					s.logger.Verbose(logger.Streams, "Broadcasting guild event %+v to user %d server %+v for guild %d", event, userID, serv, to)
					serv <- event
				} else {
					s.logger.Verbose(logger.Streams, "Not broadcasting guild event %+v to user %d server %+v for guild %d", event, userID, serv, to)
				}
			}
		}
		s.logger.Verbose(logger.Streams, "Broadcasted event %+v to guild %d", event, to)
	}()
}

// BroadcastHomeserver does a thing
func (s *StreamManager) BroadcastHomeserver(userid uint64, event *chatv1.Event) {
	s.logger.Verbose(logger.Streams, "Broadcasting homeserver event %+v to user %d", event, userid)

	go func() {
		s.RLock()
		defer s.RUnlock()

		for server := range s.userIDToServers[userid] {
			if s.serverToStreamData[server].homeserver {
				s.logger.Verbose(logger.Streams, "Broadcasting HS event %+v to user %d server %+v", event, userid, server)
				server <- event
			} else {
				s.logger.Verbose(logger.Streams, "Not broadcasting HS event %+v to user %d server %+v", event, userid, server)
			}
		}

		s.logger.Verbose(logger.Streams, "Broadcasted homeserver event %+v to user %d", event, userid)
	}()
}

// BroadcastAction does a thing
func (s *StreamManager) BroadcastAction(userid uint64, event *chatv1.Event) {
	s.logger.Verbose(logger.Streams, "Broadcasting action event %+v to user %d", event, userid)

	go func() {
		s.RLock()
		defer s.RUnlock()

		for server := range s.userIDToServers[userid] {
			if s.serverToStreamData[server].action {
				s.logger.Verbose(logger.Streams, "Broadcasting action event %+v to user %d server %+v", event, userid, server)
				server <- event
			} else {
				s.logger.Verbose(logger.Streams, "Not broadcasting action event %+v to user %d server %+v", event, userid, server)
			}
		}

		s.logger.Verbose(logger.Streams, "Broadcasted action event %+v to user %d", event, userid)
	}()
}
