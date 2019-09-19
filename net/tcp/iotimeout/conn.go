package iotimeout

import (
	"fmt"
	"github.com/dylenfu/go-libs/net/pkg/bufio"
	"math/rand"
	"net"
	"time"
)

type Connect struct {
	conn   *net.TCPConn
	rd     *bufio.Reader
	wr     *bufio.Writer
	signal chan *Msg // 用于接收外部数据， 比如说当前连接为a，当连接b有消息需要通过a发送给client时，连接a ready阻塞获取数据；如果a连接想通过b连接发送数据给客户端，需要使用b连接的push
}

func NewConnect(conn *net.TCPConn) *Connect {
	c := &Connect{}
	c.conn = conn
	c.rd = bufio.NewReader(conn)
	c.wr = bufio.NewWriter(conn)
	c.signal = make(chan *Msg, 10)
	return c
}

func (c *Connect) Close() {
	c.rd = nil
	c.wr = nil
	c.conn.Close()
	c.conn = nil
	close(c.signal)
}

func (c *Connect) Name() string {
	return c.conn.RemoteAddr().String()
}

func (c *Connect) Push(s *Msg) {
	select {
	case c.signal <- s:
	default:
		fmt.Println("channel full")
	}
}

func (c *Connect) Read(p []byte) (int, error) {
	return c.rd.Read(p)
}

func (c *Connect) Write(p []byte) (int, error) {
	return c.wr.Write(p)
}

// 30ms ~ 80ms内随机
func (c *Connect) Flush() error {
	n := rand.Intn(50)
	n += 30
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Println("flush rand time", d.String())
	return c.wr.Flush()
}

func (c *Connect) Ready() *Msg {
	return <-c.signal
}

func (c *Connect) SetWriteDeadLine(d time.Duration) error {
	t := time.Now()
	return c.conn.SetWriteDeadline(t.Add(d))
}

func (c *Connect) SetReadDeadLine(d time.Duration) error {
	t := time.Now()
	return c.conn.SetReadDeadline(t.Add(d))
}
