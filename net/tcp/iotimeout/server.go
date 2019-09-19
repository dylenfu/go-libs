package iotimeout

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	conns map[string]*Connect
}

func NewServer() *Server {
	s := &Server{}
	s.conns = make(map[string]*Connect)
	listener := s.newServerListener(host)
	go s.acceptTCP(listener)
	return s
}

func (s *Server) PushMsg(msg *Msg, conn string) {
	c := s.conns[conn]
	if c != nil {
		c.Push(msg)
	}
}

func (s *Server) newServerListener(addr string) *net.TCPListener {
	var (
		tcpaddr  *net.TCPAddr
		err      error
		listener *net.TCPListener
	)

	if tcpaddr, err = net.ResolveTCPAddr(tcp, addr); err != nil {
		log.Fatal(err)
	}
	if listener, err = net.ListenTCP(tcp, tcpaddr); err != nil {
		log.Fatal(err)
	}
	return listener
}

func (s *Server) acceptTCP(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(fmt.Sprintf("tcp connect success, local address %s to remote address %s", conn.LocalAddr().String(), conn.RemoteAddr().String()))
		}
		c := NewConnect(conn)
		s.conns[connname] = c
		go s.dispatch(c)
	}
}

func (s *Server) dispatch(c *Connect) {
	defer c.conn.Close()
	for {
		msg := c.Ready()
		bs := Encode(msg)
		if err := c.SetWriteDeadLine(writeDeadLint); err != nil {
			fmt.Println("set dead line error", err)
		}
		if _, err := c.Write(bs); err != nil {
			fmt.Println("server connection write data error", err)
			return
		}
		if err := c.Flush(); err != nil {
			fmt.Println("server connection io flush error", err)
			return
		}
	}
}
