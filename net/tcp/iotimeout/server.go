package iotimeout

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	conns map[string]*Connect
	mtx   *sync.Mutex
}

func NewServer() *Server {
	s := &Server{}
	s.conns = make(map[string]*Connect)
	s.mtx = new(sync.Mutex)
	listener := s.newServerListener(host)
	go s.acceptTCP(listener)
	return s
}

func (s *Server) PushMsg(msg *Msg, conn string) {
	for _, c := range s.conns {
		if c.conn != nil {
			c.Push(msg)
		}
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
		}

		c := NewConnect(conn)
		s.addConnect(c)
		fmt.Println("[server] tcp connect success", c.Name())
		go s.dispatch(c)
	}
}

func (s *Server) dispatch(c *Connect) {
	defer s.delConnect(c)

	for {
		msg := c.Ready()
		bs := Encode(msg)
		if err := c.SetWriteDeadLine(writeDeadLine); err != nil {
			fmt.Println("[server] set dead line error", err)
		}
		if _, err := c.Write(bs); err != nil {
			fmt.Println("[server] connection write data error", err)
			return
		}
		if err := c.Flush(); err != nil {
			fmt.Println("[server] connection io flush error", err)
			return
		}
	}
}

func (s *Server) addConnect(c *Connect) {
	s.mtx.Lock()
	s.conns[c.Name()] = c
	s.mtx.Unlock()
}

func (s *Server) delConnect(c *Connect) {
	s.mtx.Lock()
	delete(s.conns, c.Name())
	c.Close()
	s.mtx.Unlock()
}
