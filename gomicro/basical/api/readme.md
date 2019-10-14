使用micro v1.11.3版本
该示例并不需要使用consul/etcd等

1.直接运行micro api:
fukundeMacBook-Pro:api dylen$ micro api --handler=api
2019/10/14 13:52:19 Registering API Request Handler at /
2019/10/14 13:52:19 HTTP API Listening on [::]:8080
2019/10/14 13:52:19 Transport [http] Listening on [::]:64927
2019/10/14 13:52:19 Broker [http] Connected to [::]:64928
2019/10/14 13:52:19 Registry [mdns] Registering node: go.micro.api-cef84790-c397-4f8f-898d-7ef15bf45e86
127.0.0.1 - - [14/Oct/2019:13:52:24 +0800] "POST /example/foo/bar HTTP/1.1" 200 23 "" "curl/7.54.0"
127.0.0.1 - - [14/Oct/2019:13:55:23 +0800] "POST /example/foo/bar HTTP/1.1" 200 23 "" "curl/7.54.0"


2.然后运行程序 go run main.go

首先会通过micro api开启一个网关，接收来自http请求，并转换为api。
然后主程序main.go处理请求内容.

输入两个请求:
fukundeMacBook-Pro:api dylen$ curl -H 'head-1: I am a header' "http://localhost:8080/example/call?name=john"
{"message":"msg received john"}

fukundeMacBook-Pro:api dylen$ curl -H 'Content-Type: application/json' -d '{data:123}' http://localhost:8080/example/foo/bar
msg received {data:123}
