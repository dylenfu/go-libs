package main

import (
	"context"
	"fmt"
	proto "github.com/dylenfu/go-libs/gomicro/example1/greeter/proto"
	micro "github.com/micro/go-micro"
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

}
