package env

import (
	"os"
	"path"
	"log"
	"os/exec"
)

func SimplePath() {
	dir, _ := os.Getwd()
	log.Println("env\t-", "current direction:", dir)
	p := path.Dir(dir)
	log.Println("env\t-","current direction process after path.Dir:", p)
	log.Println("--------------------------------------------------------------\r\n")

	p1 := path.Dir("$GOPATH/src/github.com/Loopring/ringminer/store")
	log.Print("env\t-","path1:", p1)
	log.Println("--------------------------------------------------------------\r\n")

	gopath := os.Getenv("GOPATH")
	log.Println("env\t-","gopath:", gopath)
	log.Println("--------------------------------------------------------------\r\n")

	env := os.Environ()
	for _,v := range env {
		log.Println("env\t-","os.Environ contains:",v)
	}
	log.Println("--------------------------------------------------------------\r\n")

	lp, _ := exec.LookPath(os.Args[0])
	log.Println("env\t-", "executeable file in:", lp)
}