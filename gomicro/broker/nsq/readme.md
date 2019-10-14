使用nsq作为broker

下载nsq docker，https://nsq.io/deployment/docker.html 使用最新版本
运行 docker run --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd
然后运行
docker run --name nsqd -p 4150:4150 -p 4151:4151 \
    nsqio/nsq /nsqd \
    --broadcast-address=172.21.5.62 \
    --lookupd-tcp-address=172.21.5.62:4160

使用client发送消息，server消费消息。不同于默认的http broker，nsq broker需要在micro.NewService时创建:
micro.Broker(nsq.NewBroker(broker.Addrs([]string{NSQ_ADDR}...)))
然后将topic，server注册到subscribe。

同理，client发送消息也需要创建broker，然后，创建某个publisher，持续发送消息。


