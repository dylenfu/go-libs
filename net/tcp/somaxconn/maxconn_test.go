package somaxconn

import (
	"net"
)

const (
	serverAddr = "0.0.0.0:9090"
)

func NewServer() {
	var (
		listener *net.TCPListener
		tcpaddr  *net.TCPAddr
		err      error
	)

	if tcpaddr, err = net.ResolveTCPAddr("tcp", serverAddr); err != nil {
		panic(err)
	}
	if listener, err = net.ListenTCP("tcp", tcpaddr); err != nil {
		panic(err)
	}
	for {
		var conn *net.TCPConn
		if conn, err = listener.AcceptTCP(); err != nil {
			panic(err)
		}
		//conn.SetKeepAlive()
		serveTCP(conn)
	}
}

func serveTCP(conn *net.TCPConn) {

}
