package grpc

import (
	dc "github.com/dylenfu/go-libs/grpc/demo/client"
	ds "github.com/dylenfu/go-libs/grpc/demo/server"
	sc "github.com/dylenfu/go-libs/grpc/stream/client"
	ss "github.com/dylenfu/go-libs/grpc/stream/server"
)

func Route(sub string) {
	switch sub {
	case "simple-client":
		dc.SimpleClient()

	case "simple-server":
		ds.SimpleServer()

	case "simple-stream-client-1":
		sc.SimpleStreamClient(1)

	case "simple-stream-client-2":
		sc.SimpleStreamClient(2)

	case "simple-stream-client-3":
		sc.SimpleStreamClient(3)

	case "simple-stream-server":
		ss.SimpleStreamServer()
	}
}
