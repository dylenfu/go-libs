package main

import (
	"flag"
	"github.com/dylenfu/go-libs/base"
	"github.com/dylenfu/go-libs/cmd"
	"github.com/dylenfu/go-libs/grpc"
	"github.com/dylenfu/go-libs/http"
	"github.com/dylenfu/go-libs/inject"
	"github.com/dylenfu/go-libs/jsonrpc"
	"github.com/dylenfu/go-libs/leveldb"
	"github.com/dylenfu/go-libs/mysql"
	"github.com/dylenfu/go-libs/refers"
	"github.com/dylenfu/go-libs/serialize"
	"github.com/dylenfu/go-libs/tcp"
	"github.com/dylenfu/go-libs/toml"
	"github.com/dylenfu/go-libs/zap"
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

	case "leveldb":
		leveldb.Route(*sub)

	case "jsonrpc":
		jsonrpc.Route(*sub)

	case "inject":
		inject.Route(*sub)

	case "zap":
		zap.Route(*sub)

	case "toml":
		toml.Route(*sub)

	case "http":
		http.Route(*sub)

	case "tcp":
		tcp.Route(*sub)

	case "refers":
		refers.Route(*sub)

	case "serialize":
		serialize.Route(*sub)

	case "grpc":
		grpc.Route(*sub)

	case "mysql":
		mysql.Route(*sub)

	default:
		break
	}

}
