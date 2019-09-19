package iotimeout

import (
	"fmt"
	"io"
	"net"
	"time"
)

func NewClient() {
	var (
		remote *net.TCPAddr
		err    error
	)

	if remote, err = net.ResolveTCPAddr(tcp, host); err != nil {
		panic(err)
	}
	go connect(remote)
}

func connect(remote *net.TCPAddr) {
	var (
		conn *net.TCPConn
		err  error
	)

	for {
		time.Sleep(2 * time.Second)

		if conn, err = net.DialTCP(tcp, nil, remote); err != nil {
			panic(err)
		}

		c := NewConnect(conn)
		receive(c)
	}
}

func receive(c *Connect) {
	buf := make([]byte, bufsize)
	for {
		n, err := c.Read(buf)

		switch err {
		case io.EOF:
			fmt.Println("[client] tcp connection closed")
			c.Close()
			return

		case nil:

		default:
			fmt.Println("[client] addr error:", err)
			//c.SetReadDeadLine(10 * time.Millisecond)
		}

		bs := buf[:n]
		msg := Decode(bs)
		msg.Print()
	}
}
