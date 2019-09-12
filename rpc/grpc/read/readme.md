** 工作目录:
$GOPATH/src/github.com/dylenfu/go-libs/rpc/grpc/read

在该目录下运行build.bat保障api.pb.go的生成

该示例使用grpc，为了方便http测试，兼容http访问。

环境保障:
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

