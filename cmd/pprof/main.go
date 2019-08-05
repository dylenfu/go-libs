package main

import (
	"flag"
	. "github.com/dylenfu/go-libs/db/kv"
	"os"
	"runtime/pprof"
)

// ./pprof --cpu=cpu.prof --heap=heap.prof
// or use default path
// analysis:
// go tool pprof -svg ./pprof cpu.prof > cpu.svg
// go tool pprof -svg ./pprof heap.prof > heap.svg
var (
	cpuprofile  = flag.String("cpu", "cpu.prof", "write cpu profile to file")
	heapprofile = flag.String("heap", "heap.prof", "write mem profile to file")
	fcpu        *os.File
	fheap       *os.File
	AnalyzeHeap = NewKVCachePool(-1, USERS_SHARD_COUNT, Shard)
)

func main() {
	flag.Parse()

	openHeapProf()
	defer func() {
		closeHeapProf()
	}()

	Work()
	wait()
}

func Work() {
	for i := 1; i < 1000; i++ {
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
		AnalyzeHeap.Set(uid, user, -1)
	}
}

func openCpuProf() {
	var err error
	if fcpu, err = os.Create(*cpuprofile); err != nil {
		panic(err)
	}
	if err = pprof.StartCPUProfile(fcpu); err != nil {
		panic(err)
	}
}

func closeCpuProf() {
	pprof.StopCPUProfile()
	if fcpu != nil {
		fcpu.Close()
	}
}

func openHeapProf() {
	var err error
	if fheap, err = os.Create(*heapprofile); err != nil {
		panic(err)
	}
	if err := pprof.WriteHeapProfile(fheap); err != nil {
		panic(err)
	}
}

func closeHeapProf() {
	if fheap != nil {
		fheap.Close()
	}
}

func wait() {
	for {
		select {}
	}
}
