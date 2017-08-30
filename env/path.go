package env

import (
	"os"
	"path"
	"log"
)

func SimplePath() {
	dir, _ := os.Getwd()
	log.Println("env\t-", "current direction:", dir)
	p := path.Dir(dir)
	log.Println("env\t-","current direction process after path.Dir:", p)

	p1 := path.Dir("$GOPATH/src/github.com/Loopring/ringminer/store")
	log.Print("env\t-","path1", p1)

	gopath := os.Getenv("GOPATH")
	log.Println("env\t-","gopath", gopath)

	env := os.Environ()
	for _,v := range env {
		log.Println("env\t-","os.Environ contains",v)
	}
}