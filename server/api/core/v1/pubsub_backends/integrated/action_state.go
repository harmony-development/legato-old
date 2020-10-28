package integrated

import (
	"sync"

	corev1 "github.com/harmony-development/legato/gen/core"
)

// ActionState is the manager of action pub/sub
type ActionState struct {
	actionChannels map[corev1.CoreService_StreamActionEventsServer]chan struct{}
	actionEvents   map[_userID][]corev1.CoreService_StreamActionEventsServer
	sync.Mutex
}

// Initialize the action state
func (a *ActionState) Initialize() *ActionState {
	a.actionChannels = make(map[corev1.CoreService_StreamActionEventsServer]chan struct{})
	a.actionEvents = make(map[_userID][]corev1.CoreService_StreamActionEventsServer)
	return a
}

// Subscribe subscribes
func (a *ActionState) Subscribe(userID uint64, server corev1.CoreService_StreamActionEventsServer) chan struct{} {
	a.Lock()
	defer a.Unlock()

	go func() {
		<-server.Context().Done()
		a.Unsubscribe(userID, server)
	}()

	val, _ := a.actionEvents[_userID(userID)]
	val = append(val, server)
	a.actionEvents[_userID(userID)] = val
	a.actionChannels[server] = make(chan struct{})
	return a.actionChannels[server]
}

// Unsubscribe unsubscribes
func (a *ActionState) Unsubscribe(userID uint64, server corev1.CoreService_StreamActionEventsServer) {
	a.Lock()
	defer a.Unlock()

	val, _ := a.actionEvents[_userID(userID)]
	for idx, serv := range val {
		if serv == server {
			val[idx] = val[len(val)-1]
			val[len(val)-1] = nil
			val = val[:len(val)-1]
			break
		}
	}
	a.actionEvents[_userID(userID)] = val
	close(a.actionChannels[server])
	delete(a.actionChannels, server)
}

// Broadcast broadcasts
func (a *ActionState) Broadcast(userID uint64, action *corev1.ActionEvent) {
	val, _ := a.actionEvents[_userID(userID)]
	for _, serv := range val {
		serv.Send(action)
	}
}
