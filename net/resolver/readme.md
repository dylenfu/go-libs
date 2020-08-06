TestHostResolver 使用coreDNS，解析三个域名，主要查看net.LookupHost, splitHost等的使用.
HostsResolver 存储域名解析出来的ip地址和端口，使用unsfafe指针来操作缓存，操作方式是：
一旦超时则另外make一个对象指针存储数据，并将resolver cache指针指向新对象

coredns的使用，下载coredns docker镜像
在某个workspace下:
1.写入Corefile
. {
    log
    errors
    auto
    reload 10s
    forward . /etc/resolv.conf 
    file /root/db.ontsnip.com ontsnip.com
}

2.写入db文件db.ontsnip.com
$ORIGIN ontsnip.com.
@	3600 IN	SOA sns.dns.icann.org. dylenfu@126.com. (
				2020052604 ; serial
				60         ; refresh in seconds (1 min)
				3600       ; retry (1 hour)
				1209600    ; expire (2 weeks)
				3600       ; minimum (1 hour)
				)

	3600 IN NS a.iana-servers.net.
	3600 IN NS b.iana-servers.net.

reserve1    IN A     172.168.3.158
reserve2    IN A     172.168.3.163
reserve3    IN A     192.168.3.153

3.运行docker
#!/bin/bash

workspace=yourworkspacepath

docker stop coredns
docker rm coredns
docker run -d --name coredns \
-v=$workspace:/root/ \
-p 53:53/udp \
coredns/coredns -conf /root/Corefile

docker logs -f coredns

4.测试
#!/bin/sh

dig reserve1.ontsnip.com
dig reserve2.ontsnip.com
dig reserve3.ontsnip.com

5.注意
在linux下需要关闭系统自带的dns服务，
sudo systemctl stop systemd-resolved
在mac下，在【系统偏好】【网络】【高级】【dns】中添加本地ip(dns服务器地址)