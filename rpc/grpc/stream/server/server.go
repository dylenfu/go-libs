package server

import (
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"testing"

	pb "github.com/dylenfu/go-libs/rpc/grpcgrpc/stream/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func TestSimpleStreamServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	defer listener.Close()

	s := grpc.NewServer()
	defer s.Stop()

	pb.RegisterApiServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func (s *server) SayHello1(request *pb.HelloReq, gs pb.ApiService_SayHello1Server) error {
	for i := 0; i < 100; i++ {
		gs.Send(&pb.HelloResp{Data: "server1 " + strconv.Itoa(i) + ":" + request.Data})
	}

	return nil
}

func (s *server) SayHello2(gs pb.ApiService_SayHello2Server) error {
	var list []string

	for {
		in, err := gs.Recv()
		if err == io.EOF {
			gs.SendAndClose(&pb.HelloResp{Data: "io eof \r\n" + strings.Join(list, "\r\n")})
			return nil
		}
		if err != nil {
			log.Fatalf("failed to Recv %v", err)
			return err
		}
		list = append(list, "server2 "+in.Data)
	}

	return nil
}

func (s *server) SayHello3(gs pb.ApiService_SayHello3Server) error {
	for {
		in, err := gs.Recv()
		println(in.Data)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("failed to Recv %v", err)
			return err
		}
		gs.Send(&pb.HelloResp{Data: "server3 " + in.Data})
	}
	return nil
}
