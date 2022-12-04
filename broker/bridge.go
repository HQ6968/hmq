package broker

import (
	"go.uber.org/zap"
)

const (
	// ConnectAction mqtt connect
	ConnectAction = "connect"
	// PublishAction  mqtt publish
	PublishAction = "publish"
	// SubscribeAction  mqtt sub
	SubscribeAction = "subscribe"
	// UnsubscribeAction  mqtt sub
	UnsubscribeAction = "unsubscribe"
	// DisconnectAction mqtt disconenct
	DisconnectAction = "disconnect"
)

func (b *Broker) Publish(e *Elements) bool {
	if b.bridgeMQ != nil {
		cost, err := b.bridgeMQ.Publish(e)
		if err != nil {
			log.Error("send message to mq error.", zap.Error(err))
		}
		return cost
	}
	return false
}
