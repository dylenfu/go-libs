package main

import (
    "flag"
    "godemo/base"
    "fmt"
)

var (
    testcase = flag.String("testcase", "hi", "chose test case to execute")
)

func usage() {
    fmt.Println(`
This example creates a simple HTTP Proxy using two libp2p peers. The first peer
provides an HTTP server locally which tunnels the HTTP requests with libp2p
to a remote peer. The remote peer performs the requests and
send the sends the response back.

Usage: Start remote peer first with:   ./proxy
       Then start the local peer with: ./proxy -d <remote-peer-multiaddress>

Then you can do something like: curl -x "localhost:9900" "http://ipfs.io".
This proxies sends the request through the local peer, which proxies it to
the remote peer, which makes it and sends the response back.
`)
    flag.PrintDefaults()
}

func main() {
    flag.Usage = usage

    flag.Parse()

    switch *testcase {
    case "hi-ipfs":
        println("Hello IPFS")

    case "base-struct":
        t := &base.Tee{3, "hello"}
        println(t.GetTeeNum())
        println(t.GetTeeName())

    case "base-channel":
        base.ChannelDemo()

    case "base-reflect1":
        base.ReflectDemo1()

    case "base-reflect2":
        base.ReflectDemo2()

    case "base-reflect3":
        base.ReflectDemo3()

    case "base-reflect4":
        base.ReflectDemo4()

    case "base-reflect5":
        base.ReflectDemo5()

    case "base-reflect6":
        base.ReflectDemo6()
    }

}
