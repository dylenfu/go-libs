package main

import (
	"fmt"
	"net/http"
)

const (
	pem = "../certs/server.crt"
	key = "../certs/server.key"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":443", pem, key, nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi, This is an example1 of https service in golang!")
}
