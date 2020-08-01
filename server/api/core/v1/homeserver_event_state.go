package v1

import (
	"sync"

	corev1 "github.com/harmony-development/legato/gen/core"
)

type HomeserverEventState struct {
	homeserverEvents map[UserID][]corev1.CoreService_StreamHomeserverEventsServer
	sync.Mutex
}

var homeserverEventState = HomeserverEventState{
	homeserverEvents: map[UserID][]corev1.CoreService_StreamHomeserverEventsServer{},
}

func (h *HomeserverEventState) Subscribe(userID uint64, s corev1.CoreService_StreamHomeserverEventsServer) {
	h.Lock()
	defer h.Unlock()

	go func() {
		<-s.Context().Done()
		h.Unsubscribe(userID, s)
	}()

	if _, ok := h.homeserverEvents[UserID(userID)]; !ok {
		h.homeserverEvents[UserID(userID)] = []corev1.CoreService_StreamHomeserverEventsServer{}
	}

	h.homeserverEvents[UserID(userID)] = append(h.homeserverEvents[UserID(userID)], s)
}

func (h *HomeserverEventState) Unsubscribe(userID uint64, s corev1.CoreService_StreamHomeserverEventsServer) {
	h.Lock()
	defer h.Unlock()

	val, _ := h.homeserverEvents[UserID(userID)]
	for idx, serv := range val {
		if serv == s {
			val[idx] = val[len(val)-1]
			val[len(val)-1] = nil
			val = val[:len(val)-1]
			break
		}
	}
	h.homeserverEvents[UserID(userID)] = val
}

func (h *HomeserverEventState) Broadcast(userID uint64, e *corev1.HomeserverEvent) {
	h.Lock()
	defer h.Unlock()

	val, _ := h.homeserverEvents[UserID(userID)]
	for _, serv := range val {
		serv.Send(e)
	}
}
