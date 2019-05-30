package base

import (
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
	<-n.stop
}

func (n *Node) Close() {
	close(n.stop)
}

func TestChannel2(t *testing.T) {
	messages := make(chan string)

	msg := "hi"
	go func() {
		messages <- msg
	}()

	select {
	case msg := <-messages:
		t.Log("received message", msg)
	}
}