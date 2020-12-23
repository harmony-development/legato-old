package integrated

import (
	"sync"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
)

// ActionState is the manager of action pub/sub
type ActionState struct {
	actionChannels map[chatv1.ChatService_StreamEventsServer]chan struct{}
	actionEvents   map[_userID][]chatv1.ChatService_StreamEventsServer
	sync.Mutex
}

// Initialize the action state
func (a *ActionState) Initialize() *ActionState {
	a.actionChannels = make(map[chatv1.ChatService_StreamEventsServer]chan struct{})
	a.actionEvents = make(map[_userID][]chatv1.ChatService_StreamEventsServer)
	return a
}

// Subscribe subscribes
func (a *ActionState) Subscribe(userID uint64, server chatv1.ChatService_StreamEventsServer) chan struct{} {
	a.Lock()
	defer a.Unlock()

	go func() {
		<-server.Context().Done()
		a.Unsubscribe(userID, server)
	}()

	val := a.actionEvents[_userID(userID)]
	val = append(val, server)
	a.actionEvents[_userID(userID)] = val
	a.actionChannels[server] = make(chan struct{})
	return a.actionChannels[server]
}

// Unsubscribe unsubscribes
func (a *ActionState) Unsubscribe(userID uint64, server chatv1.ChatService_StreamEventsServer) {
	a.Lock()
	defer a.Unlock()

	val := a.actionEvents[_userID(userID)]
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
func (a *ActionState) Broadcast(userID uint64, action *chatv1.Event) {
	val, ok := a.actionEvents[_userID(userID)]
	if !ok {
		return
	}
	for _, serv := range val {
		if err := serv.Send(action); err != nil {
			println(err)
		}
	}
}
