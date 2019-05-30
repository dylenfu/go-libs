package base

import (
	"sync"
	"testing"
	"time"
)

func TestSingleton(t *testing.T) {
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
