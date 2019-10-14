package main

import (
	"context"
	proto "github.com/dylenfu/go-libs/gomicro/broker/nsq/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nsq"
)

const (
	SERVER_ID   = "go.micro.broker.nsq.srv"
	TOPIC       = "go.micro.broker.topic.nsq"
	BROEKR_ADDR = "127.0.0.1:4150"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVER_ID),
		micro.Broker(nsq.NewBroker(broker.Addrs([]string{BROEKR_ADDR}...))),
	)
	service.Init()
	sOpts := broker.NewSubscribeOptions(nsq.WithMaxInFlight(5))
	_ = micro.RegisterSubscriber(TOPIC, service.Server(), new(Sub), server.SubscriberContext(sOpts.Context))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type Sub struct{}

func (s *Sub) Process(ctx context.Context, event *proto.DemoEvent) error {
	log.Logf("Receive info: Id %d & Timestamp %d", event.Id, event.Current)
	return nil
}
