package main

import (
	"encoding/json"
	"fmt"
	proto "github.com/dylenfu/go-libs/gomicro/broker/nats/srv/proto"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"time"
)

const (
	SERVER_ID = "go.micro.cli.pubsub"
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

	b := service.Client().Options().Broker

	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go SendEvent(TOPIC, b)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func SendEvent(topic string, b broker.Broker) {
	t := time.NewTicker(5 * time.Second)

	var i int
	for _ = range t.C {
		ev := proto.Event{
			Id:        uuid.New().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("如果你看到了消息 %s, '那是因为我一直爱着你", topic),
		}

		body, _ := json.Marshal(ev)
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: body,
		}

		log.Logf("publishing %+v", ev)

		if err := b.Publish(topic, msg); err != nil {
			log.Logf("error publishing: %v", err)
		}
		i++
	}
}
