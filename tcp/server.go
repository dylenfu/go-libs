package tcp

import (
	"fmt"
	"net"
)

const host = "0.0.0.0:9090"

func server() {
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
	defer conn.Close()
	recvBuf := make([]byte, 128)
	sendBuf := []byte("server get client:" + conn.RemoteAddr().String() + " message!")
	for {
		_, err1 := conn.Read(recvBuf)
		if err1 != nil {
			println(err1)
			return
		}
		_, err2 := conn.Write(sendBuf)
		if err2 != nil {
			println(err2)
			return
		}
		fmt.Println("server recv:::::" + string(recvBuf))
	}
}
