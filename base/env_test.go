package base

import (
	"os"
	"os/exec"
	"path"
	"testing"
)

/*
2017/09/04 10:44:47 env	- current direction: /Users/fukun/projects/gohome/src/github.com/dylenfu/go-libs
2017/09/04 10:44:47 env	- current direction process after path.Dir: /Users/fukun/projects/gohome/src/github.com/dylenfu
2017/09/04 10:44:47 --------------------------------------------------------------

2017/09/04 10:44:47 env	-path1:$GOPATH/src/github.com/Loopring/ringminer
2017/09/04 10:44:47 --------------------------------------------------------------

2017/09/04 10:44:47 env	- gopath: /Users/fukun/projects/gohome
2017/09/04 10:44:47 --------------------------------------------------------------

2017/09/04 10:44:47 env	- os.Environ contains: TERM_PROGRAM=Apple_Terminal
2017/09/04 10:44:47 env	- os.Environ contains: TERM=xterm-256color
2017/09/04 10:44:47 env	- os.Environ contains: SHELL=/bin/bash
2017/09/04 10:44:47 env	- os.Environ contains: TMPDIR=/var/folders/9q/vlm5tt2n6yncq29mf4zjr8080000gn/T/
2017/09/04 10:44:47 env	- os.Environ contains: Apple_PubSub_Socket_Render=/private/tmp/com.apple.launchd.IfmsShwbna/Render
2017/09/04 10:44:47 env	- os.Environ contains: TERM_PROGRAM_VERSION=388.1.1
2017/09/04 10:44:47 env	- os.Environ contains: TERM_SESSION_ID=EDEB3A3A-EFD7-4E55-A53C-A56624B6C077
2017/09/04 10:44:47 env	- os.Environ contains: USER=fukun
2017/09/04 10:44:47 env	- os.Environ contains: SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.dHCmjMw37H/Listeners
2017/09/04 10:44:47 env	- os.Environ contains: __CF_USER_TEXT_ENCODING=0x1F5:0x19:0x34
2017/09/04 10:44:47 env	- os.Environ contains: PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/Users/fukun/softwares/go/bin:/Users/fukun/projects/gohome/bin
2017/09/04 10:44:47 env	- os.Environ contains: PWD=/Users/fukun/projects/gohome/src/github.com/dylenfu/go-libs
2017/09/04 10:44:47 env	- os.Environ contains: LANG=zh_CN.UTF-8
2017/09/04 10:44:47 env	- os.Environ contains: XPC_FLAGS=0x0
2017/09/04 10:44:47 env	- os.Environ contains: XPC_SERVICE_NAME=0
2017/09/04 10:44:47 env	- os.Environ contains: HOME=/Users/fukun
2017/09/04 10:44:47 env	- os.Environ contains: SHLVL=1
2017/09/04 10:44:47 env	- os.Environ contains: GOROOT=/Users/fukun/softwares/go
2017/09/04 10:44:47 env	- os.Environ contains: LOGNAME=fukun
2017/09/04 10:44:47 env	- os.Environ contains: GOPATH=/Users/fukun/projects/gohome
2017/09/04 10:44:47 env	- os.Environ contains: _=/Users/fukun/softwares/go/bin/go
2017/09/04 10:44:47 env	- os.Environ contains: OLDPWD=/Users/fukun/projects/gohome/src/github.com/Loopring/ringminer
2017/09/04 10:44:47 --------------------------------------------------------------

2017/09/04 10:44:47 env	- executeable file in: /var/folders/9q/vlm5tt2n6yncq29mf4zjr8080000gn/T/go-build139424569/command-line-arguments/_obj/exe/main
*/
func TestPath(t *testing.T) {
	dir, _ := os.Getwd()
	t.Log("env\t-", "current direction:", dir)
	p := path.Dir(dir)
	t.Log("env\t-", "current direction process after path.Dir:", p)
	t.Log("--------------------------------------------------------------")

	p1 := path.Dir("$GOPATH/src/github.com/Loopring/ringminer/store")
	t.Log("env\t-", "path1:", p1)
	t.Log("--------------------------------------------------------------")

	gopath := os.Getenv("GOPATH")
	t.Log("env\t-", "gopath:", gopath)
	t.Log("--------------------------------------------------------------")

	env := os.Environ()
	for _, v := range env {
		t.Log("env\t-", "os.Environ contains:", v)
	}
	t.Log("--------------------------------------------------------------")

	lp, _ := exec.LookPath(os.Args[0])
	t.Log("env\t-", "executeable file in:", lp)
}

/*
go run main.go -pkg=base -sub=simple-os-args
go test -v github.com/dylenfu/go-libs/base -run TestOsArgs
os length: 3
os.Args[0] /var/folders/9q/vlm5tt2n6yncq29mf4zjr8080000gn/T/go-build528220898/command-line-arguments/_obj/exe/main
os.Args[1] -pkg=base
os.Args[2] -sub=simple-os-args
*/
func TestOsArgs(t *testing.T) {
	t.Log("os length:", len(os.Args))
	t.Log("os.Args[0]", string(os.Args[0]))
	t.Log("os.Args[1]", string(os.Args[1]))
	t.Log("os.Args[2]", string(os.Args[2]))
}
