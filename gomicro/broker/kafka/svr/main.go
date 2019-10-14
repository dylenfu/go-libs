package main

import (
	"context"
	proto "github.com/dylenfu/go-libs/gomicro/broker/nsq/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/kafka"
)

const (
	SERVER_ID  = "go.micro.broker.kafka.srv"
	TOPIC      = "go.micro.broker.topic.kafka"
	KAFKA_ADDR = "127.0.0.1:9092"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVER_ID),
		micro.Broker(kafka.NewBroker(broker.Addrs([]string{KAFKA_ADDR}...))))

	service.Init()

	_ = micro.RegisterSubscriber(TOPIC, service.Server(), new(Sub))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type Sub struct{}

func (s *Sub) Process(ctx context.Context, event *proto.DemoEvent) error {
	log.Logf("Receive info: Id %d & Timestamp %d", event.Id, event.Current)
	return nil
}
