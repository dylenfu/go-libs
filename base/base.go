package base

import (
	"sync"
	"time"
)

type Tee struct {
	Num  uint
	Name string
}

func (tee *Tee) GetTeeNum() uint {
	return tee.Num
}

func (tee Tee) GetTeeName() string {
	return tee.Name
}

func ChannelDemo() {
	messages := make(chan string)

	msg := "hi"
	go func() {
		messages <- msg
	}()

	select {
	case msg := <-messages:
		println("received message", msg)
	}

}

func SimpleGoRoutine() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			println(i)
		}(i)
	}
	time.Sleep(10 * time.Second)
}

func SingletonDemo() {
	ton := GetInstance()
	go func() {
		for i := 0; i < 100; i++ {
			ton.do(i)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			ton.do(i)
		}
	}()

	time.Sleep(5 * time.Second)
}

var (
	once     sync.Once
	instance *singleton
)

type singleton struct{}

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (s *singleton) do(i int) {
	println(i)
}
