package iotimeout

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// 消息发送给某个client
// 手动设置server connection 发送消息延迟
// go test -v github.com/dylenfu/go-libs/net/tcp/iotimeout -run TestIoTimeOut
func TestIoTimeOut(t *testing.T) {
	s := NewServer()
	time.Sleep(1 * time.Second)
	polling(s)
	syswait()
}

func polling(s *Server) {
	for i := 0; i < 100; i++ {
		NewClient()
	}

	for i := 0; i < 100; i++ {
		// 500 ~1000ms间
		n := rand.Intn(500)
		n += 500
		time.Sleep(time.Duration(n) * time.Millisecond)
		s.PushMsg(&Msg{Data: "hello"}, connname)
	}
}

func syswait() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	for {
		sig := <-c
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP:
			fmt.Println("success msg:", msgCnt, "failed msg:", 10000-msgCnt)
			return
		case syscall.SIGHUP:
		}
	}
}
