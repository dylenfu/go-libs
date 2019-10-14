package main

import (
	"context"
	proto "github.com/dylenfu/go-libs/gomicro/broker/nsq/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
	"log"
	"time"
)

const (
	SERVER_ID  = "go.micro.broker.kafka.client"
	TOPIC      = "go.micro.broker.topic.kafka"
	KAFKA_ADDR = "127.0.0.1:9092"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVER_ID),
		micro.Broker(kafka.NewBroker(broker.Addrs([]string{KAFKA_ADDR}...))))

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
