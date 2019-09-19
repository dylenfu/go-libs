package iotimeout

import (
	"fmt"
	"io"
	"net"
	"time"
)

func NewClient() {
	var (
		local, remote *net.TCPAddr
		err           error
	)

	if local, err = net.ResolveTCPAddr(tcp, clientaddr); err != nil {
		panic(err)
	}
	if remote, err = net.ResolveTCPAddr(tcp, host); err != nil {
		panic(err)
	}
	go connect(remote, local)
}

func connect(remote, local *net.TCPAddr) {
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
	for {
		buf := make([]byte, bufsize)
		n, err := c.Read(buf)

		switch err {
		case io.EOF:
			fmt.Println("tcp connection closed")
			c.conn.Close()
			c.conn = nil
			return

		case nil:

		default:
			fmt.Println("clientaddr error:", err)
			c.SetReadDeadLine(10 * time.Millisecond)
		}

		bs := buf[:n]
		msg := Decode(bs)
		msg.Print()
	}
}
