package client

import (
	"flag"
	pb "github.com/dylenfu/go-libs/grpc/demo/rpc"
	"log"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	key    = "um.dev.tokenbank-key.pem"
	cert   = "um.dev.tokenbank-cert.pem"
	caCert = "um.dev.tokenbank-ca-cert.pem"
)

var endpoint = flag.String("endpoint", "localhost:9090", "grpc server address")

func TestSimpleClient(t *testing.T) {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	creds, err := credentials.NewClientTLSFromFile(caCert, "")
	if err != nil {
		log.Fatalf("new client tls from file %v", err)
	}

	opt := grpc.WithTransportCredentials(creds)
	conn, err := grpc.DialContext(ctx, *endpoint, opt)
	if err != nil {
		log.Fatalf("dail context %v", err)
	}
	defer conn.Close()

	client := pb.NewApiServiceClient(conn)

	info, err := client.SayHello(ctx, &pb.HelloReq{Data: "test"})
	if err != nil {
		log.Fatalf("say hello %v", err)
	}

	log.Print("greeting is %s", info.Data)
}
