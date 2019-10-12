package main

import (
	"context"
	"fmt"
	proto "github.com/dylenfu/go-libs/gomicro/example1/greeter/proto"
	micro "github.com/micro/go-micro"
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

	service.Init()

	_ = proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
