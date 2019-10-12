package main

import (
	"context"
	"fmt"
	proto "github.com/dylenfu/go-libs/gomicro/example1/greeter/proto"
	cli2 "github.com/micro/cli"
	"github.com/micro/go-micro"
	"time"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, res *proto.HelloResponse) error {
	res.Greeting = "Hello" + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)

	service.Init(micro.Action(func(i *cli2.Context) {

	}))

	_ = proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	go Producer(service, "greeter.topic.pubsub.1")
	go Producer(service, "greeter.topic.pubsub.2")

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func Producer(service micro.Service, topic string) {
	pub := micro.NewPublisher(topic, service.Client())

	for {
		time.Sleep(5 * time.Second)
		event := &proto.HelloEvent{Name: "your daddy", Age: 32}
		if err := pub.Publish(context.TODO(), event); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(topic, event)
		}
	}
}
