package sense

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

type user struct {
	uid  int
	name string
}

type room struct {
	users map[int]*user
	wch   chan interface{}
	rch   chan *user
}

func getinstance() *room {
	r := &room{}
	r.users = make(map[int]*user)
	r.wch = make(chan interface{})
	r.rch = make(chan *user)

	// wch这个通道，允许同时接收写入数据及查询条件，如此才能保证读写数据的互斥性，实现互斥锁的功能
	go func() {
		for x := range r.wch {
			switch n := x.(type) {
			case *user:
				r.users[n.uid] = n
			case int:
				r.rch <- r.users[n]
			default:
				panic("data type invalid")
			}
		}
	}()
	return r
}

func (r *room) safeWrite(user *user) {
	r.wch <- user
}

func (r *room) safeRead(uid int) *user {
	r.wch <- uid
	return <-r.rch
}

func TestChannelInsteadMutex(t *testing.T) {
	r := getinstance()
	wg := sync.WaitGroup{}
	for i := 1; i < 10000; i++ {
		wg.Add(1)
		go func(uid int) {
			r.safeWrite(&user{uid, "test" + strconv.Itoa(uid)})
			wg.Done()
		}(i)
	}
	for i := 1; i < 10000; i++ {
		wg.Add(1)
		go func(uid int) {
			fmt.Println(r.safeRead(uid))
		}(i)
	}
	wg.Wait()
}
