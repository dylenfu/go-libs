
注意:我们在代码中有用到,
namespace = go.micro.evt
topic=user
serverid=user2

那么，该server的全称应该是go.micro.evt.user2.
在开启网关的时候,需要注册网关:
micro api --handler=event --namespace=go.micro.evt

然后再运行主程序.运行后发行，Process3会导致报错，因为其入口参数缺失context&interface，
也就是说，sub消费应该是一个实现了func(context.Context, interface{}) error的接口。

除了Process3会导致报错以外，当我们使用curl:
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/login
会发现Process1&2都有打印信息，但是process没有打印信息，
这说明私有方法是不能被执行的。

连续使用curl:
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/login
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/login1
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/login100
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/register122
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/register122
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/register122/match
发现，只要是在topic user下，上述方式都是被允许执行的

2019/10/14 14:16:08 Transport [http] Listening on [::]:49682
2019/10/14 14:16:08 Broker [http] Connected to [::]:49683
2019/10/14 14:16:08 Registry [mdns] Registering node: user2-294b6190-86a3-44c2-a565-d7bbdc316c30
2019/10/14 14:16:08 Subscribing user2-294b6190-86a3-44c2-a565-d7bbdc316c30 to topic: go.micro.evt.user
2019/10/14 14:18:15 public Process1 received event:login
2019/10/14 14:18:15 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:18:15 public Process2 received event:login
2019/10/14 14:18:15 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:20:35 public Process2 received event:login1
2019/10/14 14:20:35 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:20:35 public Process1 received event:login1
2019/10/14 14:20:35 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:20:45 public Process2 received event:login100
2019/10/14 14:20:45 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:20:45 public Process1 received event:login100
2019/10/14 14:20:45 public Process2 received data:{"message": "Hello, Micro中国"}
^[[A^[[B2019/10/14 14:21:16 public Process2 received event:register122
2019/10/14 14:21:16 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:21:16 public Process1 received event:register122
2019/10/14 14:21:16 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:21:32 public Process2 received event:register122
2019/10/14 14:21:32 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:21:32 public Process1 received event:register122
2019/10/14 14:21:32 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:28:14 public Process2 received event:register122.match
2019/10/14 14:28:14 public Process1 received event:register122.match
2019/10/14 14:28:14 public Process2 received data:{"message": "Hello, Micro中国"}
2019/10/14 14:28:14 public Process2 received data:{"message": "Hello, Micro中国"}

但是一旦user变成user2就不行了，这是因为topic发生了变化。