package main

import (
	"context"
	"fmt"
	proto "github.com/dylenfu/go-libs/gomicro/example1/greeter/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter.client"),
	)
	service.Init()

	// Create a new greater client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call greeter
	req := &proto.HelloRequest{Name: "dylenfu"}
	if res, err := greeter.Hello(context.TODO(), req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	Consumer1(service)
	Consumer2(service)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// 消费有两种模式，只要是某个interface就可以
func Consumer1(service micro.Service) {
	_ = micro.RegisterSubscriber("greeter.topic.pubsub.1", service.Server(), new(sub))
}

// 这里最后一个参数有没有都可以，使用queue.pubsub是传入一个option
func Consumer2(service micro.Service) {
	_ = micro.RegisterSubscriber("greeter.topic.pubsub.2", service.Server(), SubEv, server.SubscriberQueue("queue.pubsub"))
}

type sub struct{}

func (s *sub) Process(ctx context.Context, event *proto.HelloEvent) error {
	fmt.Println("::::1", event.Name, event.Age)
	return nil
}

func SubEv(ctx context.Context, event *proto.HelloEvent) error {
	fmt.Println("::::2", event.Name, event.Age)
	return nil
}
