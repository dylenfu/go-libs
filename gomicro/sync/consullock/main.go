package consullock

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/sync/lock"
	"github.com/micro/go-micro/sync/lock/consul"
	"time"
)

// 模拟多服务竞争分布式锁
func main() {
	node := "localhost:8500"

	nodes := lock.Nodes(node)
	resourceId := "id"

	go func() {
		lc := consul.NewLock(nodes)
		log.Log("goroutine1 is getting sync lock......")
		if err := lc.Acquire(resourceId); err != nil {
			log.Log("[ERR] goroutine1 get lock failed!")
			return
		}

		log.Log("goroutine1 get lock success, waiting 1 second......")
		time.Sleep(1 * time.Second)

		log.Log("goroutine1 release lock")
		if err := lc.Release(resourceId); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		lc := consul.NewLock(nodes)
		log.Log("goroutine2 is getting sync lock......")
		if err := lc.Acquire(resourceId); err != nil {
			log.Log("[ERR] goroutine2 get lock failed!")
			return
		}

		log.Log("goroutine2 get lock success, waiting 1 second......")
		time.Sleep(1 * time.Second)

		log.Log("goroutine2 release lock")
		if err := lc.Release(resourceId); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(5 * time.Second)
}
