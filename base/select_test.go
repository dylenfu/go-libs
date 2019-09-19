package base

import (
	"fmt"
	"testing"
	"time"
)

// go test -v github.com/dylenfu/go-libs/base -run TestBufferredChannel
func TestBufferredChannel(t *testing.T) {
	c := new(Channel)
	c.signal = make(chan int, 10)
	for i := 0; i < 20; i++ {
		go c.push(i)
	}
	go c.dispatch()

	time.Sleep(5 * time.Second)
}

type Channel struct {
	signal chan int
}

func (c *Channel) push(p int) (err error) {
	select {
	case c.signal <- p:
	default:
		fmt.Println("channel full")
	}
	return
}

func (c *Channel) dispatch() {
	for {
		i := <-c.signal
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}

// go test -v github.com/dylenfu/go-libs/base -run TestMultipleFlush
func TestMultipleFlush(t *testing.T) {
	ts := time.Now().UnixNano()
	time.Sleep(10 * time.Millisecond)
	te := time.Now().UnixNano()
	t.Log((te - ts) / int64(time.Millisecond))
}
