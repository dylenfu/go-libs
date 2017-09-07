package ref

import (
	"reflect"
	"net/http"
	"fmt"
)

type fakeHandler struct{}
func (frw *fakeHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func handle(pattern string, handler interface{}) {
	handlerInterface := reflect.TypeOf(new(http.Handler)).Elem()
	handlerFunction  := reflect.TypeOf(new(http.HandlerFunc)).Elem()
	t := reflect.TypeOf(handler)
	if t.Implements(handlerInterface) {fmt.Println("http.Handler")}
	//http.HandlerFunc is a different type than
	// func(http.ResponseWriter, *http.Request), but we can do
	// var hf HandlerFunc = func(http.ResponseWriter, *http.Request){}
	if t.AssignableTo(handlerFunction) {fmt.Println("http.HandleFunc")}
}

func f(http.ResponseWriter, *http.Request) {}

func SimpleReflect() {
	handle("",&fakeHandler{})
	handle("",f)
}
