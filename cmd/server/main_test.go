package main

import (
	"fmt"
	"github.com/fhmq/hmq/broker"
	"go.uber.org/zap"
	"testing"
)

type mockBridge struct{}

func (m mockBridge) Publish(e *broker.Elements) (bool, error) {
	fmt.Println("mockBridge.Publish", e)
	return true, nil
}

type mockAuth struct{}

func (m mockAuth) CheckConnect(clientID, username, password string) bool {
	fmt.Println("mockAuth.CheckConnect", clientID, username, password)
	return true
}

func TestRun(t *testing.T) {
	b, err := broker.NewBroker(broker.DefaultConfig,
		broker.WithBridge(&mockBridge{}),
		broker.WithAuth(&mockAuth{}),
	)
	if err != nil {
		log.Fatal("New Broker error: ", zap.Error(err))
	}
	b.Start()

	s := waitForSignal()
	log.Info("signal received, broker closed.", zap.Any("signal", s))
}
