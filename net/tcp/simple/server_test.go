package simple

import (
	"fmt"
	"net"
	"testing"
)

const host = "0.0.0.0:9090"

func TestServer(t *testing.T) {
	address, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		panic(err)
		return
	}

	listener, err := net.ListenTCP("tcp4", address)
	if err != nil {
		panic(err)
	}

	acceptTCP(listener)
}

func acceptTCP(listener *net.TCPListener) {
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			println(err.Error())
			continue
		}
		handleServerConnection(conn)
	}
}

func handleServerConnection(conn *net.TCPConn) {
	var err error

	defer conn.Close()
	recvBuf := make([]byte, 128)
	sendBuf := []byte("server get client:" + conn.RemoteAddr().String() + " message!")
	for {
		if _, err = conn.Read(recvBuf); err != nil {
			println(err)
			return
		}
		if _, err = conn.Write(sendBuf); err != nil {
			println(err)
			return
		}
		fmt.Println("server recv:::::" + string(recvBuf))
	}
}
