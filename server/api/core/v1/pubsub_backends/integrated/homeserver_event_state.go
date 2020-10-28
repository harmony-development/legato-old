package integrated

import (
	"sync"

	corev1 "github.com/harmony-development/legato/gen/core"
)

// HomeserverEventState ...
type HomeserverEventState struct {
	homeserverChannels map[corev1.CoreService_StreamHomeserverEventsServer]chan struct{}
	homeserverEvents   map[_userID][]corev1.CoreService_StreamHomeserverEventsServer
	sync.Mutex
}

// Initialize the homeserver event state
func (h *HomeserverEventState) Initialize() *HomeserverEventState {
	h.homeserverChannels = make(map[corev1.CoreService_StreamHomeserverEventsServer]chan struct{})
	h.homeserverEvents = make(map[_userID][]corev1.CoreService_StreamHomeserverEventsServer)
	return h
}

// Subscribe ...
func (h *HomeserverEventState) Subscribe(userID uint64, s corev1.CoreService_StreamHomeserverEventsServer) chan struct{} {
	h.Lock()
	defer h.Unlock()

	go func() {
		<-s.Context().Done()
		h.Unsubscribe(userID, s)
	}()

	if _, ok := h.homeserverEvents[_userID(userID)]; !ok {
		h.homeserverEvents[_userID(userID)] = []corev1.CoreService_StreamHomeserverEventsServer{}
	}

	h.homeserverChannels[s] = make(chan struct{})
	h.homeserverEvents[_userID(userID)] = append(h.homeserverEvents[_userID(userID)], s)
	return h.homeserverChannels[s]
}

// Unsubscribe ...
func (h *HomeserverEventState) Unsubscribe(userID uint64, s corev1.CoreService_StreamHomeserverEventsServer) {
	h.Lock()
	defer h.Unlock()

	val, _ := h.homeserverEvents[_userID(userID)]
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
func (h *HomeserverEventState) Broadcast(userID uint64, e *corev1.HomeserverEvent) {
	h.Lock()
	defer h.Unlock()

	val, _ := h.homeserverEvents[_userID(userID)]
	for _, serv := range val {
		serv.Send(e)
	}
}
