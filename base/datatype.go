package base

import (
	"fmt"
	"sync/atomic"
	"time"
)

func SimpleStringType() {
	bytes1 := []byte{'a', 'b', 'c', 'd', 'e'}
	str1 := string(bytes1)
	bytes2 := []byte(str1)
	var bytes3 []byte = []byte{'a', 'b', 'c', 'd', 'e'}

	println("str1", str1)
	println("bytes2.size", len(bytes2))
	for k, v := range bytes3 {
		println(k, byte(v))
		fmt.Println(k, string(v))
	}
}

// 最小的原子操作，atomic.load保证加载的数据不被cpu读写
func SimpleAtomic() {
	countDown := int64(300)

	var (
		lastTimes int64
		diff      int64
		nowCount  int64
		timer     = int64(5)
	)

	for {
		nowCount = atomic.LoadInt64(&countDown)
		diff = nowCount - lastTimes
		lastTimes = nowCount
		fmt.Println(fmt.Sprintf("%s down:%d down/s:%d", time.Now().Format("2006-01-02 15:04:05"), nowCount, diff/timer))
		time.Sleep(time.Duration(timer) * time.Second)
	}
}
