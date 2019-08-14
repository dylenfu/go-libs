package base

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleChannel1(t *testing.T) {
	var n Node
	n.Start()
	go n.Wait()
	time.Sleep(5 * time.Second)
	n.Close()
}

type Node struct {
	ID   int
	Desc string
	stop chan struct{}
}

func (n *Node) Start() {
	n.ID = 1001
	n.Desc = "node start and stop test"
	n.stop = make(chan struct{})
}

func (n *Node) Wait() {
	fmt.Println(<-n.stop)
}

func (n *Node) Close() {
	close(n.stop)
}

func TestSimpleChannel2(t *testing.T) {
	messages := make(chan string)

	msg := "hi"
	go func() {
		messages <- msg
	}()

	select {
	case m := <-messages:
		t.Log("received message", m)
	}
}

// select用于在多个通道中进行选择，他会一直阻塞，直到通道数据准备就绪
// 下面的两个例子:
// DeadLock1,因为通道没有数据，执行时报错DeadLock,
// DeadLock2,在select中添加default，那么select会执行default语句，也就不会造成死锁
// DeadLock3,不给通道添加缓冲，流入后阻塞
// DeadLock4,缓冲通道，流入后不阻塞
// DeadLock5,通道流入与流出不在同一个gorutine，不阻塞，DeadLock6同理
func TestSelectDeadLock1(t *testing.T) {
	c := make(chan string)
	select {
	case m := <-c:
		t.Log(m)
	}
}

func TestSelectDeadLock2(t *testing.T) {
	c := make(chan string)
	select {
	case m := <-c:
		t.Log(m)
	default:
	}
}

func TestSelectDeadLock3(t *testing.T) {
	c := make(chan string)
	c <- "hi"
	m := <-c
	t.Log(m)
}

func TestSelectDeadLock4(t *testing.T) {
	c := make(chan string, 1)
	c <- "hi"
	m := <-c
	t.Log(m)
}

func TestSelectDeadLock5(t *testing.T) {
	c := make(chan string)
	go func() { c <- "hi" }()
	m := <-c
	t.Log(m)
}

func TestSelectDeadLock6(t *testing.T) {
	c := make(chan string)
	go func() { c <- "hi" }()
	select {
	case m := <-c:
		t.Log(m)
	}
}

// 测试select执行顺序
// select case如果有计算语句，计算顺序为先左后右，先上后下
// select case在执行时，如果多个case都有可能则随机执行
func TestSelectExecutionOrder(t *testing.T) {
	c := make(chan int, 5)

	for i := 0; i < cap(c); i++ {
		select {
		case c <- 1:
		case c <- 2:
		case c <- 3:
		}
	}

	for i := 0; i < cap(c); i++ {
		t.Log(<-c)
	}
}

//// 使用channel 替代 mutex
//type Op interface {
//	Key() string
//	Value() interface{}
//}
//
//func server() chan<- Op {
//	ch := make(chan Op)
//	m := map[string]string{}
//	go func() {
//		for op := range ch {
//			k, v := op.Key(), op.Value()
//			if s, ok := v.(string); ok {
//				m[k] = s
//				continue
//			}
//			v.(chan string) <- m[k]
//		}
//	}()
//	return ch
//}
//
//type s struct {
//	key   string
//	value string
//}
//
//func (x *s) Key() string        { return x.key }
//func (x *s) Value() interface{} { return x.value }
//
//func setOp(key, value string) Op {
//	return &s{key, value}
//}
//
//type g struct {
//	key   string
//	value chan string
//}
//
//func (x *g) Key() string        { return x.key }
//func (x *g) Value() interface{} { return x.value }
//
//func getOp(key string) Op {
//	return &g{key, make(chan string)}
//}
//
//func TestChannelInMap(t *testing.T) {
//	srv := server()
//	//srv <- setOp("key", "value")
//	//srv <- setOp("k3", "v3")
//	//srv <- setOp("k2", "v2")
//	//op := getOp("k3")
//	//srv <- op
//	//fmt.Println(<-(op.(*g).value))
//	for i := 0; i < 10000; i++ {
//		k := strconv.Itoa(i)
//		v := "test_" + k
//		go func(k, v string) {
//			srv <- setOp(k, v)
//		}(k, v)
//	}
//
//	for i := 0; i < 10001; i++ {
//		go func(k string) {
//			op := getOp(k)
//			srv <- op
//			fmt.Println(k, <-op.(*g).value)
//		}(strconv.Itoa(i))
//	}
//
//	time.Sleep(5 * time.Second)
//}
//
//
/////----------------测试单项写通道在函数返回后是否还有写入的能力
//func TestWriteOnlyChannel(t *testing.T) {
//	var data  = map[string]int{}
//	srv := writeOnly(data)
//	for i := 0; i < 100; i++ {
//		srv <- &kv{"test_" + strconv.Itoa(i), i}
//	}
//	for k, v := range data {
//		t.Log(k, v)
//	}
//}
//
//type kv struct {
//	k string
//	v int
//}
//
//func writeOnly(data map[string]int) chan <- *kv {
//	ch := make(chan *kv)
//	go func() {
//		for x := range ch {
//			data[x.k] = x.v
//		}
//	}()
//	return ch
//}
//
//func readOnly(data map[string]int, key string) <- chan *kv {
//	ch := make(chan *kv)
//	go func() {
//		for x := range ch {
//			x.v
//		}
//	}()
//	return ch
//}
