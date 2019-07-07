package base

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleChannel1(t *testing.T) {
	var n Node
	n.Start()
	go n.Wait()
	time.Sleep(5 * time.Second)
	n.Close()
}

type Node struct {
	ID   int
	Desc string
	stop chan struct{}
}

func (n *Node) Start() {
	n.ID = 1001
	n.Desc = "node start and stop test"
	n.stop = make(chan struct{})
}

func (n *Node) Wait() {
	fmt.Println(<-n.stop)
}

func (n *Node) Close() {
	close(n.stop)
}

func TestSimpleChannel2(t *testing.T) {
	messages := make(chan string)

	msg := "hi"
	go func() {
		messages <- msg
	}()

	select {
	case m := <-messages:
		t.Log("received message", m)
	}
}

// select用于在多个通道中进行选择，他会一直阻塞，直到通道数据准备就绪
// 下面的两个例子:
// DeadLock1,因为通道没有数据，执行时报错DeadLock,
// DeadLock2,在select中添加default，那么select会执行default语句，也就不会造成死锁
// DeadLock3,不给通道添加缓冲，流入后阻塞
// DeadLock4,缓冲通道，流入后不阻塞
// DeadLock5,通道流入与流出不在同一个gorutine，不阻塞，DeadLock6同理
func TestSelectDeadLock1(t *testing.T) {
	c := make(chan string)
	select {
		case m := <-c:
			t.Log(m)
	}
}

func TestSelectDeadLock2(t *testing.T) {
	c := make(chan string)
	select {
		case m := <- c:
			t.Log(m)
		default:
	}
}

func TestSelectDeadLock3(t *testing.T) {
	c := make(chan string)
	c <- "hi"
	m := <- c
	t.Log(m)
}

func TestSelectDeadLock4(t *testing.T) {
	c := make(chan string, 1)
	c <- "hi"
	m:= <-c
	t.Log(m)
}

func TestSelectDeadLock5(t *testing.T) {
	c := make(chan string)
	go func(){c <- "hi"}()
	m:= <-c
	t.Log(m)
}

func TestSelectDeadLock6(t *testing.T) {
	c := make(chan string)
	go func(){c <- "hi"}()
	select {
		case m := <- c:
			t.Log(m)
	}
}

// 测试select执行顺序
// select case如果有计算语句，计算顺序为先左后右，先上后下
// select case在执行时，如果多个case都有可能则随机执行
func TestSelectExecutionOrder(t *testing.T) {
	c := make(chan int, 5)

	for i:=0; i < cap(c); i++ {
		select {
			case c <- 1:
			case c <- 2:
			case c <- 3:
		}
	}

	for i:=0; i < cap(c); i++ {
		t.Log(<-c)
	}
}
