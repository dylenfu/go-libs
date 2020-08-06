package resolver

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// TestHostResolver 使用coreDNS，解析三个域名，主要查看net.LookupHost, splitHost等的使用.
// HostsResolver 存储域名解析出来的ip地址和端口，使用unsfafe指针来操作缓存，操作方式是：
// 一旦超时则另外make一个对象指针存储数据，并将resolver cache指针指向新对象
// go test -v github.com/dylenfu/go-libs/net/resolver -run TestHostResolver
func TestHostResolver(t *testing.T) {
	hosts := []string{
		"reserve1.ontsnip.com:20336",
		"reserve2.ontsnip.com:20336",
		"reserve3.ontsnip.com:20336",
	}
	resolver, invalids := NewHostsResolver(hosts)
	fmt.Printf("invalids %v\r\n", invalids)

	timer := time.NewTimer(0)
	for {
		select {
		case <-timer.C:
			for _, address := range resolver.GetHostAddrs() {
				fmt.Printf("address cache %s\r\n", address)
			}
			timer.Stop()
			timer.Reset(10 * time.Second)
		default:
		}
	}
}

// HostsResolver host resovler with cache
type HostsResolver struct {
	hosts [][2]string

	lock  sync.Mutex     // avoid concurrent cache reflesh
	cache unsafe.Pointer // atomic pointer to HostsCache, avoid read&write data race
}

// HostsCache save addresses
type HostsCache struct {
	refleshTime time.Time
	addrs       []string
}

// NewHostsResolver new resolver instance
func NewHostsResolver(hosts []string) (*HostsResolver, []string) {
	resolver := &HostsResolver{}
	var invalids []string
	for _, n := range hosts {
		host, port, e := net.SplitHostPort(n)
		if e != nil {
			invalids = append(invalids, n)
			continue
		}
		resolver.hosts = append(resolver.hosts, [2]string{host, port})
	}

	return resolver, invalids
}

// GetHostAddrs get host addrs
func (s *HostsResolver) GetHostAddrs() []string {
	// fast path test
	cached := (*HostsCache)(atomic.LoadPointer(&s.cache))
	if cached != nil && cached.refleshTime.Add(time.Minute*10).After(time.Now()) && len(cached.addrs) != 0 {
		return cached.addrs
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	cached = (*HostsCache)(s.cache)
	if cached != nil && cached.refleshTime.Add(time.Minute*10).After(time.Now()) && len(cached.addrs) != 0 {
		return cached.addrs
	}

	cache := make([]string, 0, len(s.hosts))
	for _, n := range s.hosts {
		host, port := n[0], n[1]
		// 本机dns服务
		ns, err := net.LookupHost(host)
		if err != nil || len(ns) == 0 {
			continue
		}

		cache = append(cache, net.JoinHostPort(ns[0], port))
	}

	atomic.StorePointer(&s.cache, unsafe.Pointer(&HostsCache{refleshTime: time.Now(), addrs: cache}))

	return cache
}
