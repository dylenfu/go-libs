package main

import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/util/log"
	"time"
)

const (
	TOPIC = "mu.micro.book.topic.payment.done"
)

func main() {
	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("broker init error:%v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker connect error:%v", err)
	}
	go pub()
	go sub()
	<-time.After(time.Second * 20)
}

func pub() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for range ticker.C {
		msg := &broker.Message{
			Header: map[string]string{"id": fmt.Sprintf("%d", i)},
			Body:   []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		log.Infof(broker.String())
		if err := broker.Publish(TOPIC, msg); err != nil {
			log.Infof("[pub] msg publish failed:%v", err)
		} else {
			log.Infof("[pub] msg %s publish success", string(msg.Body))
		}
		i++
	}
}

func sub() {
	_, err := broker.Subscribe(TOPIC, func(p broker.Event) error {
		log.Infof("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil {
		log.Log(err.Error())
	}
}
