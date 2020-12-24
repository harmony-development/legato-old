package integrated

import (
	"sync"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/harmony-development/legato/server/logger"
)

// HomeserverEventState ...
type HomeserverEventState struct {
	homeserverChannels map[chatv1.ChatService_StreamEventsServer]chan struct{}
	homeserverEvents   map[_userID][]chatv1.ChatService_StreamEventsServer
	Logger             logger.ILogger
	sync.Mutex
}

// Initialize the homeserver event state
func (h *HomeserverEventState) Initialize(l logger.ILogger) *HomeserverEventState {
	h.Logger = l
	h.homeserverChannels = make(map[chatv1.ChatService_StreamEventsServer]chan struct{})
	h.homeserverEvents = make(map[_userID][]chatv1.ChatService_StreamEventsServer)
	return h
}

// Subscribe ...
func (h *HomeserverEventState) Subscribe(userID uint64, s chatv1.ChatService_StreamEventsServer) chan struct{} {
	h.Lock()
	defer h.Unlock()

	go func() {
		<-s.Context().Done()
		h.Unsubscribe(userID, s)
	}()

	if _, ok := h.homeserverEvents[_userID(userID)]; !ok {
		h.homeserverEvents[_userID(userID)] = []chatv1.ChatService_StreamEventsServer{}
	}

	h.homeserverChannels[s] = make(chan struct{})
	h.homeserverEvents[_userID(userID)] = append(h.homeserverEvents[_userID(userID)], s)
	return h.homeserverChannels[s]
}

// Unsubscribe ...
func (h *HomeserverEventState) Unsubscribe(userID uint64, s chatv1.ChatService_StreamEventsServer) {
	h.Lock()
	defer h.Unlock()

	val, ok := h.homeserverEvents[_userID(userID)]
	_ = ok
	for idx, serv := range val {
		if serv == s {
			val[idx] = val[len(val)-1]
			val[len(val)-1] = nil
			val = val[:len(val)-1]
			break
		}
	}
	close(h.homeserverChannels[s])
	delete(h.homeserverChannels, s)
	h.homeserverEvents[_userID(userID)] = val
}

// Broadcast ...
func (h *HomeserverEventState) Broadcast(userID uint64, e *chatv1.Event) {
	h.Lock()
	defer h.Unlock()

	val, ok := h.homeserverEvents[_userID(userID)]
	_ = ok
	for _, serv := range val {
		if err := serv.Send(e); err != nil {
			println(err)
		}
	}
}
