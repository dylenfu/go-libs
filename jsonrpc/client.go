package jsonrpc

import (
	"net"
	"time"
	"log"
	"net/rpc/jsonrpc"
)

func NewClient1() {
	// 新建连接
	conn, err := net.DialTimeout("tcp", "localhost:8888", 1 * time.Second)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	// 创建jsonrpc客户端
	client := jsonrpc.NewClient(conn)

	// 远程服务器返回的对象
	var req ReqMessage
	client.Call("GetName", 1,&req)
	log.Println("client\t-", "call GetName method", req)
}
