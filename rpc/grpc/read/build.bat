## checkout https://github.com/grpc-ecosystem/grpc-gateway

## generate api.pb.go
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  api.proto

## generate api.pb.gw.go
protoc -I/usr/local/include -I. \
  -I${GOPATH}/src \
  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/  \
  --grpc-gateway_out=logtostderr=true:. \
  api.proto

## generate swagger.json
protoc -I/usr/local/include -I. \
  -I${GOPATH}/src \
  -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/  \
  --swagger_out=logtostderr=true:. \
  api.proto
