broker示例:

broker是一种分布式模式，解耦客户端和服务端，服务端将服务注册到broker，通过接口暴露的方式允许客户端接入，
客户端将请求发送到broker，broker转发请求到服务端，并将请求结果返回给客户端。

这种方式的好处是服务端可以动态地进行更新、删除和添加，对于客户端来说，这些都是透明的。

该示例提供了默认的http broker。展示了broker的相关操作。
package go-micro/broker提供了一个DefaultBroker，

Init动作对broker进行option初始化，
Connect连接相关服务
Publish发布消息
Subscribe消费消息
