package main

// ref: https://github.com/micro-in-cn/tutorials/tree/master/examples/basic-practices/micro-api/api
// 使用API类型的服务，需要用到专门定义的原型服务 github.com/micro/go-micro/api/proto/api.proto
// 运行api网关 micro api --handler=api

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/dylenfu/go-libs/gomicro/basical/api/proto"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	"strings"
)

const (
	SERVER_ID           = "go.micro.api.example"
	POST                = "POST"
	GET                 = "GET"
	STATUS_CODE_SUCCESS = 200
)

func main() {
	service := micro.NewService(
		micro.Name(SERVER_ID),
	)

	service.Init()

	_ = pb.RegisterExampleHandler(service.Server(), new(Example))
	_ = pb.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type Example struct{}

func (e *Example) Call(ctx context.Context, req *api.Request, res *api.Response) error {
	var (
		name *api.Pair
		ok   bool
	)

	log.Log("Example.Call receive request")

	// validate params
	if name, ok = req.Get["name"]; !ok || len(name.Values) == 0 {
		return fmt.Errorf("params invalid")
	}

	// show request header
	for k, v := range req.Header {
		log.Log("request header info, k:", k, ", v:", v)
	}

	b, _ := json.Marshal(map[string]string{
		"message": fmt.Sprintf("msg received %s", strings.Join(name.Values, " ")),
	})
	res.Body = string(b)
	res.StatusCode = STATUS_CODE_SUCCESS

	return nil
}

type Foo struct{}

func (f *Foo) Bar(ctx context.Context, req *api.Request, res *api.Response) error {
	var (
		content *api.Pair
		ok      bool
	)

	log.Log("Foo.Bar received request")

	if req.Method != POST {
		return errors.BadRequest(SERVER_ID, "method require post")
	}
	if content, ok = req.Header["Content-Type"]; !ok || len(content.Values) == 0 {
		return errors.BadRequest(SERVER_ID, "content type required")
	}
	if content.Values[0] != "application/json" {
		return errors.BadRequest(SERVER_ID, "expect application/json")
	}

	var body = map[string]interface{}{}
	_ = json.Unmarshal([]byte(req.Body), &body)

	res.Body = fmt.Sprintf("msg received %s", string([]byte(req.Body)))
	res.StatusCode = STATUS_CODE_SUCCESS

	return nil
}
