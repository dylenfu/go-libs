package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

const (
	SERVER_ID = "go.micro.srv.pubsub"
	TOPIC     = "go.micro.pubsub.topic.event"
	VERSION   = "latest"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(SERVER_ID),
		micro.Version(VERSION),
	)

	// Initialise service
	service.Init()

	broker := service.Server().Options().Broker

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	if _, err := broker.Subscribe(TOPIC, Handler); err != nil {
		log.Fatalf("broker.Subscribe error: %v", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func Handler(event broker.Event) error {
	log.Logf("[sub] 收到消息，请查收: %v, %v", event.Message().Header, string(event.Message().Body))
	return nil
}
