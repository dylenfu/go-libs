package main

import (
	"flag"
	"github.com/dylenfu/go-libs/cmd"
	"github.com/dylenfu/go-libs/mysql"
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
	
	case "mysql":
		mysql.Route(*sub)

	default:
		break
	}

}
