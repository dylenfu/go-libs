package jsonrpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
	"time"
)

var (
	req  ReqMessage
	resp RespMessage
)

func newClient() *rpc.Client {
	// 新建连接
	conn, err := net.DialTimeout("tcp", "localhost:"+port, 1*time.Second)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 创建jsonrpc客户端
	client := jsonrpc.NewClient(conn)
	return client
}

// aync call
func TestAyncCall(t *testing.T) {
	client := newClient()
	defer client.Close()

	// 同步请求
	client.Call("ServerHandle.GetName", 1, &req)
	log.Println("client\t-", "call GetName method", req)
}

// sync call
func TestSyncCall(t *testing.T) {
	client := newClient()
	defer client.Close()

	req.Id = 5
	req.Name = "ha"

	// 异步请求
	syncCall := client.Go("ServerHandle.SetName", &req, &resp, nil)
	go func() {
		_ = syncCall.Done
		//log.Println("client\t-", "call SetName", reply)
	}()
	time.Sleep(1 * time.Second)

	log.Println("client\t-", "call SetName response ", resp)
}
