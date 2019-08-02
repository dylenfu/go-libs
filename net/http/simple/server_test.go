package simple

import (
	"fmt"
	"net/http"
	"testing"
)

type mutex struct{}

func TestSimpleHttpServer(t *testing.T) {
	http.HandleFunc("/hello", handleTest)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}

//////////////////////////////////////////////////
// 这里要注意的是返回给调用方数据时需要使用Fprintf
// 此外，解析get的数据可以用parseForm
//////////////////////////////////////////////////
func handleTest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data, _ := r.Form["data"]
	fmt.Fprintf(w, "hello "+data[0])
}
