package base

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestStringType(t *testing.T) {
	bytes1 := []byte{'a', 'b', 'c', 'd', 'e'}
	str1 := string(bytes1)
	bytes2 := []byte(str1)
	var bytes3 []byte = []byte{'a', 'b', 'c', 'd', 'e'}

	t.Logf("str1:%s", str1)
	t.Logf("bytes2.size:%d", len(bytes2))
	for k, v := range bytes3 {
		//t.Log(k, byte(v))
		t.Log(k, string(v))
	}
}

// 最小的原子操作，atomic.load保证加载的数据不被cpu读写
func TestAtomic(t *testing.T) {
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
		t.Log(fmt.Sprintf("%s down:%d down/s:%d", time.Now().Format("2006-01-02 15:04:05"), nowCount, diff/timer))
		time.Sleep(time.Duration(timer) * time.Second)
	}
}
