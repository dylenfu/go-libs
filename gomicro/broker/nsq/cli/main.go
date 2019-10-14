package main

import (
	"context"
	proto "github.com/dylenfu/go-libs/gomicro/broker/nsq/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/nsq"
	"log"
	"time"
)

const (
	SERVER_ID = "go.micro.broker.nsq.client"
	TOPIC     = "go.micro.broker.topic.nsq"
	NSQ_ADDR  = "127.0.0.1:4150"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVER_ID),
		micro.Broker(nsq.NewBroker(broker.Addrs([]string{NSQ_ADDR}...))))

	service.Init()

	pub := micro.NewPublisher(TOPIC, service.Client())
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			_ = pub.Publish(context.TODO(), &proto.DemoEvent{
				Id:      int32(i),
				Current: time.Now().Unix(),
			})
		}
	}()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
