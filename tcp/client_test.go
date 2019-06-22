package tcp

import (
	"fmt"
	"net"
	"testing"
)

func Testclient(t *testing.T) {
	address, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		panic(fmt.Sprintf("net.ResolveTCPAddr(\"tcp4\", \"%s\") error(%v)", address.String(), err.Error()))
	}
	conn, err := net.DialTCP("tcp4", nil, address)
	if err != nil {
		panic(fmt.Sprintf("net.DailTCP(\"tcp4\", \"%s\") error(%v)", address.String(), err.Error()))
	}
	handleConnection(conn)
}

func handleConnection(conn *net.TCPConn) {
	sendBuf := []byte(conn.LocalAddr().String() + " get server message!")
	recvBuf := make([]byte, 128)
	for {
		_, err1 := conn.Write(sendBuf)
		if err1 != nil {
			println(fmt.Sprintf("conn.Write(\"%s\") error(%v)", string(sendBuf), err1))
			return
		}
		_, err2 := conn.Read(recvBuf)
		if err2 != nil {
			println(fmt.Sprintf("conn.Read(\"%s\") error(%v)", string(sendBuf), err2))
			return
		}
		fmt.Println("client recv:::::" + string(recvBuf))
	}
	defer conn.Close()
}
