package jsonrpc

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"net"
)

type ServerHandle struct {}

func (s *ServerHandle) GetName(id int, req *ReqMessage) error {
	log.Println("server\t-", "receive GetName call")
	req.Id = id
	req.Name = "dylenfu"

	return nil
}

func (s *ServerHandle) SetName(req *ReqMessage, resp *RespMessage) error {
	log.Println("server\t-", "recive SaveName call, RpcObj:", req)
	resp.Ok = true
	resp.Id = req.Id
	resp.Msg = "存储成功"

	return nil
}

func NewServer1() {
	// 新建服务器
	server := rpc.NewServer()

	// 开始监听
	listener,err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("server\t-", "listen error:", err.Error())
	}
	defer listener.Close()
	log.Println("server\t-", "start listen on 8888")

	// 新建处理器
	sh := &ServerHandle{}
	// 注册处理器
	server.RegisterName("ServerHandle", sh)

	// 等待并处理连接
	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}