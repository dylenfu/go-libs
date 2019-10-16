package sourcecode

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

// 协程被阻塞，运行时，死锁检测机制检测到没有唤醒机制，会panic
func TestEmptySelect(t *testing.T) {
	select {}
}

//因为selectgo会将cases顺序打乱，所以执行结果有三种可能，但是default可能性最大
func TestRandSelectCases(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 1
		time.Sleep(5 * time.Second)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	default:
		fmt.Println("default")
	}
	fmt.Println("main exit.")
}

// 通道关闭的情况下，chan1&chan2都成为nil，nil通道的<-操作会给出一个0
func TestCloseChannel(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()

	go func() {
		close(chan2)
	}()

	select {
	case x := <-chan1:
		fmt.Println("chan1 ready.", x)
	case x := <-chan2:
		fmt.Println("chan2 ready.", x)
	}

	fmt.Println("main exit.")
}

// func selectgo源码根据cas0创建[]scase时代码如下
// cas1 := (*[1 << 16]scase)(unsafe.Pointer(cas0))
// 籍此建立了一个长度为65536的数组
func TestSelectCases(t *testing.T) {
	type scase struct {
		//c           *hchan        // chan
		elem        unsafe.Pointer // data element
		kind        uint16
		pc          uintptr // race pc (for race detector / msan)
		releasetime int64
	}

	cas0 := &scase{}
	cas1 := (*[1 << 16]scase)(unsafe.Pointer(cas0))

	addr0 := uintptr(unsafe.Pointer(cas0))
	addr1 := uintptr(unsafe.Pointer(&cas1[0]))
	if addr0 != addr1 {
		t.Fatal("addr0 != addr1")
	}
	size := uintptr(32) // 在内存中scase占用32字节

	addr13 := uintptr(unsafe.Pointer(&cas1[3]))
	if addr13 != uintptr(addr0+size*3) {
		t.Fatal("addr13 != uintptr(addr0 + size * 3)", addr0, addr13)
	}
	addr1n := uintptr(unsafe.Pointer(&cas1[65535]))
	if addr1n != uintptr(addr0+size*65535) {
		t.Fatal("addr1n != uintptr(addr0 + size * 65535)", addr0, addr1n)
	}
}
