该示例使用consul实现分布式锁.

首先，运行consul
consul agent -dev -bind 172.**.**.62 -server -bootstrap -data-dir /yourpath

该示例中使用两个goroutine争抢同一个资源resourceId，
