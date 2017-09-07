package jsonrest

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"reflect"
)

////////////////////////////////////////////////////////////////////////////////////////
//
// 使用appsimple建立app
//
////////////////////////////////////////////////////////////////////////////////////////
func Hello() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"body":"hello json rest"})
	}))

	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

////////////////////////////////////////////////////////////////////////////////////////
//
// 使用makerouter建立app
//
////////////////////////////////////////////////////////////////////////////////////////
func SimpleRoute() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	route,err := rest.MakeRouter(
		rest.Get("/test1", testHandle1),
		rest.Get("/test2", testHandle2),
	)
	if err != nil {
		panic(err)
	}

	api.SetApp(route)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func testHandle1(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"body": "simple route handle test1"})
}

func testHandle2(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"body": "simple route handle test2"})
}

////////////////////////////////////////////////////////////////////////////////////////
//
// 通过请求字段及参数映射到相关函数
//
////////////////////////////////////////////////////////////////////////////////////////

func Reflect() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	handles := makeHandlers()
	route,err := rest.MakeRouter(handles...)

	if err != nil {
		panic(err)
	}

	api.SetApp(route)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func makeHandlers() []*rest.Route {
	var list []*rest.Route

	type api struct {
		name string
		method string
	}

	funclist := []*api{}
	funclist = append(funclist, &api{name:"Test1", method: "GET"})
	funclist = append(funclist, &api{name:"Test2", method: "GET"})

	controller := reflect.ValueOf(&orderbook{})
	for _, v := range funclist {
		route := new(rest.Route)
		method := v.method
		name := v.name

		route.HttpMethod = method
		route.PathExp = "/" + name

		route.Func = func(w rest.ResponseWriter, r *rest.Request) {
			method := controller.MethodByName(name)
			data := method.Call([]reflect.Value{})
			w.WriteJson(map[string]string{"body": data[0].String()})
		}

		list = append(list, route)
	}

	return list
}

type orderbook struct {}

func (ob *orderbook) Test1() string  {
	return "it is test 11111"
}

func (ob *orderbook) Test2() string {
	return "it is test 22222"
}
