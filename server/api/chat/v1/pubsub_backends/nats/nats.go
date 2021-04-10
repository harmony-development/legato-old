package nats

import (
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/nats-io/nats.go"
)

type NatsManager struct {
}

func New() *NatsManager {
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("mu", func(msg *nats.Msg) {

	})

	return &NatsManager{}
}

func (n *NatsManager) RegisterClient(userID uint64, s chan *chatv1.Event, done chan struct{}) {

}

func (n *NatsManager) AddGuildSubscription(s chan *chatv1.Event, to uint64) {

}

func (n *NatsManager) AddHomeserverSubscription(s chan *chatv1.Event) {

}

func (n *NatsManager) AddActionSubscription(s chan *chatv1.Event) {

}

func (n *NatsManager) RemoveGuildSubscription(userID, guildID uint64) {

}

func (n *NatsManager) BroadcastGuild(to uint64, event *chatv1.Event) {

}

func (n *NatsManager) BroadcastHomeserver(userid uint64, event *chatv1.Event) {

}

func (n *NatsManager) BroadcastAction(userid uint64, event *chatv1.Event) {

}
