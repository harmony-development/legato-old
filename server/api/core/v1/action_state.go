package v1

import (
	"sync"

	corev1 "github.com/harmony-development/legato/gen/core"
)

type ActionState struct {
	actionChannels map[corev1.CoreService_StreamActionEventsServer]chan struct{}
	actionEvents   map[UserID][]corev1.CoreService_StreamActionEventsServer
	sync.Mutex
}

func (a *ActionState) AddAction(userID uint64, server corev1.CoreService_StreamActionEventsServer) chan struct{} {
	a.Lock()
	defer a.Unlock()

	go func() {
		<-server.Context().Done()
		a.RemoveAction(userID, server)
	}()

	val, _ := a.actionEvents[UserID(userID)]
	val = append(val, server)
	a.actionEvents[UserID(userID)] = val
	a.actionChannels[server] = make(chan struct{})
	return a.actionChannels[server]
}

func (a *ActionState) RemoveAction(userID uint64, server corev1.CoreService_StreamActionEventsServer) {
	a.Lock()
	defer a.Unlock()

	val, _ := a.actionEvents[UserID(userID)]
	for idx, serv := range val {
		if serv == server {
			val[idx] = val[len(val)-1]
			val[len(val)-1] = nil
			val = val[:len(val)-1]
			break
		}
	}
	a.actionEvents[UserID(userID)] = val
	close(a.actionChannels[server])
	delete(a.actionChannels, server)
}

func (a *ActionState) BroadcastAction(userId uint64, action *corev1.ActionEvent) {
	val, _ := a.actionEvents[UserID(userId)]
	for _, serv := range val {
		serv.Send(action)
	}
}
