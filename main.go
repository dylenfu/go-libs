package main

import (
	"flag"
	"github.com/dylenfu/go-libs/base"
	"github.com/dylenfu/go-libs/jsonrpc"
	"github.com/dylenfu/go-libs/leveldb"
	"github.com/dylenfu/go-libs/zap"
	"github.com/dylenfu/go-libs/inject"
	"github.com/dylenfu/go-libs/cmd"
	"github.com/dylenfu/go-libs/env"
)

var (
	testcase = flag.String("testcase", "hi", "chose test case to execute")
)

func main() {
	flagToRun()
	//cliToRun()
}

func cliToRun() {
	cmd.SimpleCli()
}

func flagToRun() {
	flag.Parse()

	switch *testcase {
	case "hi-ipfs":
		println("Hello IPFS")
		break

	case "base-struct":
		t := &base.Tee{3, "hello"}
		println(t.GetTeeNum())
		println(t.GetTeeName())
		break

	case "base-channel":
		base.ChannelDemo()
		break

	case "simple-channel-wait":
		base.SimpleChannelDemo()
		break

	case "base-interface":
		base.InterfaceDemo()
		break

	case "env-simple-path":
		env.SimplePath()
		break

	case "reflect1":
		base.ReflectDemo1()
		break

	case "reflect2":
		base.ReflectDemo2()
		break

	case "reflect3":
		base.ReflectDemo3()
		break

	case "reflect4":
		base.ReflectDemo4()
		break

	case "reflect5":
		base.ReflectDemo5()
		break

	case "reflect6":
		base.ReflectDemo6()
		break

	case "jsonrpc-server1":
		jsonrpc.NewServer1()
		break

	case "jsonrpc-aync-call":
		jsonrpc.AyncCall()
		break

	case "jsonrpc-sync-call":
		jsonrpc.SyncCall()
		break

	case "leveldb-simple-put-get":
		leveldb.SimplePutAndGet()
		break

	case "leveldb-simple-batch":
		leveldb.SimpleBatch()
		break

	case "leveldb-batch-load":
		leveldb.SimpleBatchLoad()
		break

	case "leveldb-get-property":
		leveldb.SimpleGetProperty()
		break

	case "leveldb-get-snapshot":
		leveldb.SimpleGetSnapshot()
		break

	case "leveldb-db-iterator":
		leveldb.SimpleNewDBIterator()
		break

	case "leveldb-iterator-seek":
		leveldb.SimpleDBIteratorSeek()
		break

	case "leveldb-iterator-prefix":
		leveldb.SimpleIteratorWithPrefix()
		break

	case "leveldb-filter":
		leveldb.SimpleFilter()
		break

	case "zap-quick-start":
		zap.SimpleZapLogger()
		break

	case "zap-simple-save":
		zap.SimpleSavingZapLogger()
		break

	case "zap-multi-save":
		zap.MultipleSavingZapLogger()
		break

	case "zap-logger-print":
		zap.SimpleLoggerAndPrint()
		break

	case "inject-start":
		inject.SimpleInject()
		break

	case "inject-mine":
		inject.RewriteFaceBookInjectDemo()
		break

	case "inject-interface":
		inject.InjectInterface()
		break

	default:
		break
	}

}
