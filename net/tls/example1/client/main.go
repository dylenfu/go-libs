package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url        = "https://localhost:443"
	caCertPath = "/Users/dylen/workspace/gohome/src/github.com/dylenfu/go-libs/net/tls/example1/certs/ca.crt"
)

func main() {
	//defaultRequest()
	//ignoreVerifyCert()
	verifyCert()
}

// request without tls
// 默认客户端会校验服务端证书，这里我们的证书domain是www.random.com
// return "Get https://localhost:443: x509: certificate is valid for www.random.com, not localhost"
func defaultRequest() {
	res, err := http.Get(url)
	handlerResponse(res, err)
}

// request without verify server pem
// 客户端不再校验服务端证书
// success
func ignoreVerifyCert() {
	trans := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{Transport: trans}
	res, err := client.Get(url)
	handlerResponse(res, err)
}

// request verify server pem
// 客户端主动校验证书
// 类似于浏览器，需要预先加载ca证书
// 返回: Hi, This is an example1 of https service in golang!
func verifyCert() {
	pool := x509.NewCertPool()
	cert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	pool.AppendCertsFromPEM(cert)

	trans := http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := http.Client{Transport: &trans}

	res, err := client.Get(url)
	handlerResponse(res, err)
}

func handlerResponse(res *http.Response, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(body))
	}
}
