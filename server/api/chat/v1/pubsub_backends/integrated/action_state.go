package integrated

import (
	"sync"

	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/harmony-development/legato/server/logger"
)

// ActionState is the manager of action pub/sub
type ActionState struct {
	actionChannels map[chatv1.ChatService_StreamEventsServer]chan struct{}
	actionEvents   map[_userID][]chatv1.ChatService_StreamEventsServer
	Logger         logger.ILogger
	sync.Mutex
}

// Initialize the action state
func (a *ActionState) Initialize(l logger.ILogger) *ActionState {
	a.Logger = l
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
		a.Logger.Debug(logger.Streams, "Disconnecting", userID, server, "from actions stream")
		a.Unsubscribe(userID, server)
	}()

	a.Logger.Debug(logger.Streams, "Connecting & subscribing", userID, server, "to actions stream")
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

	a.Logger.Debug(logger.Streams, "Unsubscribing", userID, server, "from actions stream")

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
		a.Logger.Debug(logger.Streams, "Broadcast actions to a user without any actions streams", userID)
		return
	}
	for _, serv := range val {
		a.Logger.Debug(logger.Streams, "Broadcasting action event to", userID, serv)
		if err := serv.Send(action); err != nil {
			println(err)
		}
	}
}
