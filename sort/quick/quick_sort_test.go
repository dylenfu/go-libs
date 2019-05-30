package quick

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// cmd is:
// go test -v github.com/dylenfu/algorithm/sort/quick -run Test_Qsort
func TestQsort(t *testing.T) {
	values := []int{3, 7, 8, 5, 2, 1, 9, 5, 4}
	Qsort(values)
	fmt.Println(values)
}

func TestSyncMap(t *testing.T) {
	m1 := &sync.Map{}
	m1.Store("t2", 2)
	m1.Store("t3", 3)
	if value, ok :=m1.Load("t2"); ok {
		fmt.Println(value)
	}

	t.Log("println")
	fmt.Println("fmt.Println")
}

func TestConcurrentMap(t *testing.T) {
	cm := &sync.Map{}
	cm.Store("testing1", 1)
	cm.Store("testing2", 2)
	cm.Delete("testing1")
	if value, exist := cm.Load("testing1"); exist {
		t.Log(value)
	} else {
		t.Error("key not exist")
	}
	cm.Range(func (k, v interface{}) bool {
		t.Logf("key:%s value:%d", k, v)
		return true
	})
}

func TestWaitGroup(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i:=1; i< 5; i++ {
		wg.Add(1)
		go func(j int){
			fmt.Printf("test %d", j)
		}(i)
		wg.Done()
	}
	wg.Wait()
}

func TestMultilGorutine(t *testing.T) {
	data := make(chan int)
	go func(d chan int) {
		for {
			select {
			case m :=<- d:
				t.Logf("data is %d", m)
			}
		}
	}(data)
	time.Sleep(5 * time.Second)
	data <- 3
}

func TestTimerAfter(t *testing.T) {
	c := 1
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Printf("timer counting:%d\r\n", c)
			c = c + 1
		}
	}
}

func TestTimerAfterFunc(t *testing.T) {
	f := func() {
		fmt.Printf("count:%d\r\n", 3)
	}
	time.AfterFunc(1 * time.Second, f)
	time.Sleep(5 * time.Second)
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <- ticker.C:
			fmt.Println("test")
		}
	}
}

func TestTicker2(t *testing.T) {
	c := time.Tick(1 * time.Second)
	for x := range c {
		fmt.Println(x)
	}
}

func TestTicker3(t *testing.T) {
	c  := time.Tick(1 * time.Millisecond)
	for now := range c {
		fmt.Println(now)
	}
}