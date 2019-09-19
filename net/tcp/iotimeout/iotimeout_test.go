package iotimeout

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// 消息发送给某个client
// 手动设置server connection 发送消息延迟
func TestIoTimeOut(t *testing.T) {
	s := NewServer()
	time.Sleep(1 * time.Second)
	NewClient()
	polling(s)
	syswait()
}

func polling(s *Server) {
	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)
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
			return
		case syscall.SIGHUP:
		}
	}
}
