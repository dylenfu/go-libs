package simple

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
	var err error

	defer conn.Close()

	sendBuf := []byte(conn.LocalAddr().String() + " get server message!")
	recvBuf := make([]byte, 128)
	for {
		if _, err = conn.Write(sendBuf); err != nil {
			println(fmt.Sprintf("conn.Write(\"%s\") error(%v)", string(sendBuf), err))
			return
		}
		if _, err = conn.Read(recvBuf); err != nil {
			println(fmt.Sprintf("conn.Read(\"%s\") error(%v)", string(sendBuf), err))
			return
		}
		fmt.Println("client recv:::::" + string(recvBuf))
	}
}
