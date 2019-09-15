package read

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"
)

const (
	key    = "um.dev.tokenbank-key.pem"
	cert   = "um.dev.tokenbank-cert.pem"
	caCert = "um.dev.tokenbank-ca-cert.pem"

	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24

	port      = 9998
	proxyPort = 9999
)

type helloService struct{}

func registerServer(t *testing.T) {
	var (
		lis        net.Listener
		credential credentials.TransportCredentials
		err        error
	)

	// port
	sport := portString(port)

	// params
	// timeout = 1 * time.Second
	idleTimeout := 60 * time.Second
	maxLifeTime := 2 * time.Hour
	forceCloseWait := 20 * time.Second
	keepAliveInterval := 60 * time.Second
	keepAliveTimeout := 20 * time.Second

	// set server option
	keepParams := keepalive.ServerParameters{
		MaxConnectionIdle:     idleTimeout,
		MaxConnectionAgeGrace: forceCloseWait,
		Time:                  keepAliveInterval,
		Timeout:               keepAliveTimeout,
		MaxConnectionAge:      maxLifeTime,
	}
	option := grpc.KeepaliveParams(keepParams)

	// set tls credential
	if credential, err = credentials.NewServerTLSFromFile(cert, key); err != nil {
		t.Fatal(err)
	}
	credOpt := grpc.Creds(credential)

	// create grpc server
	server := grpc.NewServer(option, credOpt)

	// register server
	service := &helloService{}
	RegisterApiServiceServer(server, service)

	// start server
	if lis, err = net.Listen("tcp", sport); err != nil {
		t.Fatal(err)
	}
	go func() {
		if err = server.Serve(lis); err != nil {
			t.Fatal(err)
		}
	}()
}

func registerProxy(t *testing.T, ctx context.Context) {
	var (
		credential credentials.TransportCredentials
		err        error
	)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	endpoint := portString(proxyPort)
	grpcSrv := portTarget(port)

	// set credentail
	if credential, err = credentials.NewClientTLSFromFile(caCert, ""); err != nil {
		t.Fatal(err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(credential)}

	// set mutex
	mux := runtime.NewServeMux()

	// register proxy
	if err = RegisterApiServiceHandlerFromEndpoint(ctx, mux, grpcSrv, opts); err != nil {
		t.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err = http.ListenAndServe(endpoint, mux); err != nil {
		t.Fatal(err)
	}
}

func (s *helloService) SayHello(ctx context.Context, req *HelloRequest) (res *HelloResponse, err error) {
	res = &HelloResponse{Name: "dylenfu", Age: 32}
	return
}

func generateClient(t *testing.T, ctx context.Context, port int) ApiServiceClient {
	var (
		conn       *grpc.ClientConn
		credential credentials.TransportCredentials
		err        error
	)

	// set credentail
	if credential, err = credentials.NewClientTLSFromFile(caCert, ""); err != nil {
		t.Fatal(err)
	}

	// grpc options
	grpcKeepAliveTime := 10 * time.Second
	grpcKeepAliveTimeout := 3 * time.Second
	grpcBackoffMaxDelay := 3 * time.Second
	grpcMaxSendMsgSize := 1 << 24
	grpcMaxCallMsgSize := 1 << 24

	// set options
	// notice: grpc.WithInsecure & WithTransportCredentials can not been use at the same time
	params := []grpc.DialOption{
		//grpc.WithInsecure(),
		grpc.WithInitialWindowSize(grpcInitialWindowSize),
		grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
		grpc.WithBackoffMaxDelay(grpcBackoffMaxDelay),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                grpcKeepAliveTime,
			Timeout:             grpcKeepAliveTimeout,
			PermitWithoutStream: true,
		}),
		grpc.WithTransportCredentials(credential),
	}
	target := portTarget(port)

	if conn, err = grpc.DialContext(ctx, target, params...); err != nil {
		t.Fatal(err)
	}

	return NewApiServiceClient(conn)
}

func portString(port int) string { return ":" + strconv.Itoa(port) }
func portTarget(port int) string { return "localhost" + portString(port) }

/////////////////////////////////////////////////////////////////////////////
//
// 该测试主要用于debug及阅读代码，观察grpc的创建以及http2中的各个细节
//
//
/////////////////////////////////////////////////////////////////////////////

func TestSayHello(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Hour))
	defer cancel()

	registerServer(t)
	registerProxy(t, ctx)
	client := generateClient(t, ctx, port)
	req := HelloRequest{Id: 1}
	if res, err := client.SayHello(ctx, &req); err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
	//syswait()
}

func syswait() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGSTOP, syscall.SIGTERM)
	for {
		sig := <-c
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGQUIT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
