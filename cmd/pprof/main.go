package main

import (
	"flag"
	"fmt"
	. "github.com/dylenfu/go-libs/db/kv"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

// ./pprof --cpu=cpu.prof --heap=heap.prof
// or use default path
// analysis:
// go tool pprof -svg ./pprof cpu.prof > cpu.svg
// go tool pprof -svg ./pprof heap.prof > heap.svg
// go tool pprof  -text  -inuse_objects  ./pprof  heap.txt
var (
	cpuprofile  = flag.String("cpu", "", "write cpu profile to file")
	heapprofile = flag.String("heap", "", "write mem profile to file")
)

func main() {
	flag.Parse()

	var (
		fcpu *os.File
		err  error
	)

	if *cpuprofile != "" {
		fmt.Println("start cpu profile writing...")
		if fcpu, err = os.Create(*cpuprofile); err != nil {
			panic(err)
		}
		if err = pprof.StartCPUProfile(fcpu); err != nil {
			panic(err)
		}
		defer func() {
			pprof.StopCPUProfile()
			fcpu.Close()
		}()
	}

	TestKVSet()

	// waiting for signal to stop
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func TestKVSet() {
	var (
		fheap *os.File
		heap  = NewKVCachePool(-1, USERS_SHARD_COUNT, Shard)
		err   error
	)

	if *heapprofile == "" {
		panic("heap file can not be empty")
	}
	if fheap, err = os.Create(*heapprofile); err != nil {
		panic(err)
	}

	for i := 1; i < 10000; i++ {
		uid := 150000 + i
		user := &UserData{
			Uid:      int32(uid),
			Nickname: "ddd",
			Gender:   0,
			Verified: 1,
			Portrait: "/ssssss",
			Exp:      1,
			Level:    1,
		}
		heap.Set(uid, user, -1)
		fmt.Println(heap.Get(uid))
	}

	if err = pprof.WriteHeapProfile(fheap); err != nil {
		panic(err)
	}
	fheap.Close()
}
