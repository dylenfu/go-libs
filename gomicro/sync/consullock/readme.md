该示例使用consul实现分布式锁.

首先，运行consul
consul agent -dev -bind 172.**.**.62 -server -bootstrap -data-dir /yourpath

fukundeMacBook-Pro:go-libs dylen$ lsof -i:8500
COMMAND  PID  USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
consul  4206 dylen   14u  IPv4 0xc1c292a0a3402439      0t0  TCP localhost:fmtp (LISTEN)
fukundeMacBook-Pro:go-libs dylen$ ps -ef|grep consul
  501  4206  4205   0  4:56下午 ttys002    0:14.58 consul agent -dev -bind 172.21.5.62 -server -bootstrap -data-dir /Users/dylen/software/consul/data
  501  4329   864   0  5:11下午 ttys008    0:00.00 grep consul
fukundeMacBook-Pro:go-libs dylen$

首先，需要有一个consul服务，使用consul默认端口
node := "localhost:8500"
nodes := lock.Nodes(node)

根据nodes建立一个分布式锁，
该示例中使用两个goroutine争抢同一个资源resourceId，用完后释放
效果大致如下:

2019/10/14 16:57:12 goroutine2 is getting sync lock......
2019/10/14 16:57:12 goroutine1 is getting sync lock......
2019/10/14 16:57:12 goroutine1 get lock success, waiting 1 second......
2019/10/14 16:57:13 goroutine1 release lock
2019/10/14 16:57:13 goroutine2 get lock success, waiting 1 second......
2019/10/14 16:57:14 goroutine2 release lock


注意，这里为了使用consul，我们将micro版本切换到了1.8，当前版本1.11用的是etcd。
