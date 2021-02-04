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

	s.serverToStreamData[srv] = streamData{
		userID: userID,
		guilds: make(map[uint64]struct{}),
	}

	servs, ok := s.userIDToServers[userID]
	unused(ok)
	if servs == nil {
		s.userIDToServers[userID] = make(map[chan *chatv1.Event]struct{})
		servs = s.userIDToServers[userID]
	}
	servs[srv] = struct{}{}

	s.userIDToServers[userID] = servs

	go func() {
		<-done

		s.Lock()
		defer s.Unlock()

		delete(s.serverToStreamData, srv)
		delete(s.userIDToServers[userID], srv)

		if len(s.userIDToServers[userID]) == 0 {
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

	strct := s.serverToStreamData[srv]
	strct.homeserver = true
	s.serverToStreamData[srv] = strct
}

// AddActionSubscription adds a subscription
func (s *StreamManager) AddActionSubscription(srv chan *chatv1.Event) {
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
					serv <- event
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
				server <- event
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
				server <- event
			}
		}
	}()
}
