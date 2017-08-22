package main

import (
    "flag"
    "godemo/base"
    "fmt"
    "godemo/jsonrpc"
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
        break

    case "base-struct":
        t := &base.Tee{3, "hello"}
        println(t.GetTeeNum())
        println(t.GetTeeName())
        break

    case "base-channel":
        base.ChannelDemo()
        break

    case "base-interface":
        base.InterfaceDemo()
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

    default:
        break
    }

}
