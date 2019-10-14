package main

import (
	"context"
	"fmt"
	proto "github.com/dylenfu/go-libs/gomicro/basical/meta/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/api"
	rapi "github.com/micro/go-micro/api/handler/api"
	"github.com/micro/go-micro/api/handler/rpc"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
)

const (
	SERVER_ID = "go.micro.api.example"
)

// Endpoint is a mapping between an RPC method and HTTP endpoint
func main() {
	service := micro.NewService(
		micro.Name(SERVER_ID),
	)

	service.Init()

	_ = proto.RegisterExampleHandler(service.Server(), new(Example), api.WithEndpoint(&api.Endpoint{
		Name:        "Example.Call",
		Description: "",
		Path:        []string{"/example"},
		Method:      []string{"POST"},
		Handler:     rpc.Handler,
	}))

	_ = proto.RegisterFooHandler(service.Server(), new(Foo), api.WithEndpoint(&api.Endpoint{
		Name:    "Foo.Bar",
		Path:    []string{"/foo/bar"},
		Method:  []string{"POST"},
		Handler: rapi.Handler,
	}))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type Example struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, res *proto.CallResponse) error {
	log.Log("Meta Example.Call received request")

	if len(req.Name) == 0 {
		return errors.BadRequest(SERVER_ID, "request name empty")
	}

	res.Message = fmt.Sprintf("Meta received %s's message", req.Name)
	return nil
}

type Foo struct{}

func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, res *proto.EmptyResponse) error {
	log.Log("Meta Foo.Bar received request")
	return nil
}
