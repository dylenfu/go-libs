package main

import (
	"context"
	"github.com/micro/go-micro"
	proto "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/util/log"
)

const (
	NAMESPACE = "go.micro.evt"
	SERVER_ID = "user2"
	TOPIC     = NAMESPACE + ".user"
)

func main() {
	service := micro.NewService(
		micro.Name(SERVER_ID),
	)

	service.Init()

	if err := micro.RegisterSubscriber(TOPIC, service.Server(), new(Event)); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type Event struct{}

func (e *Event) Process1(ctx context.Context, event *proto.Event) error {
	log.Log("public Process1 received event:", event.Name)
	log.Log("public Process2 received data:", event.Data)
	return nil
}

func (e *Event) Process2(ctx context.Context, event *proto.Event) error {
	log.Log("public Process2 received event:", event.Name)
	log.Log("public Process2 received data:", event.Data)
	return nil
}

// 如果开启该方法会导致报错:
// 2019/10/14 14:10:34 subscriber Event.Process3 takes wrong number of args: 1
// required signature func(context.Context, interface{}) error
/*
func (e *Event) Process3() error {
	log.Log("public Process3 received event, nothing extracting")
	return nil
}
*/

// 非公有方法不会被调用
func (e *Event) process(ctx context.Context, event *proto.Event) error {
	log.Log("private process received event:", event.Name)
	log.Log("private process received data:", event.Data)
	return nil
}
