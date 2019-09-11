package server

import (
	pb "github.com/dylenfu/go-libs/rpc/grpc/demo/rpc"
	"log"
	"net"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9090"
	cert = "um.dev.tokenbank-cert.pem"
	key  = "um.dev.tokenbank-key.pem"
)

type server struct{}

func TestSimpleServer(t *testing.T) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatalf("failed to server tls from file %v", err)
	}

	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)

	pb.RegisterApiServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{"hello" + request.Data}, nil
}
