package main

import (
	"flag"
	"github.com/dylenfu/go-libs/base"
	"github.com/dylenfu/go-libs/jsonrpc"
	"github.com/dylenfu/go-libs/leveldb"
	"github.com/dylenfu/go-libs/zap"
	"github.com/dylenfu/go-libs/inject"
	"github.com/dylenfu/go-libs/cmd"
	"github.com/dylenfu/go-libs/toml"
	"github.com/dylenfu/go-libs/http"
	"github.com/dylenfu/go-libs/tcp"
	"github.com/dylenfu/go-libs/refers"
)

var (
	pkg = flag.String("pkg", "base", "chose package to use")
	sub = flag.String("sub", "hi", "chose sub case")
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

	switch *pkg {

	case "base":
		base.Route(*sub)
		break
	case "leveldb":
		leveldb.Route(*sub)
		break
	case "jsonrpc":
		jsonrpc.Route(*sub)
		break
	case "inject":
		inject.Route(*sub)
		break
	case "zap":
		zap.Route(*sub)
		break
	case "toml":
		toml.Route(*sub)
		break
	case "http":
		http.Route(*sub)
		break
	case "tcp":
		tcp.Route(*sub)
		break
	case "refers":
		refers.Route(*sub)
		break
	default:
		break
	}

}
