package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	SERVER_ID = "go.micro.api.example"
	POST      = "POST"
	GET       = "GET"
)

func main() {
	service := web.NewService(
		web.Name(SERVER_ID))

	service.HandleFunc("/example/call", exampleCall)
	service.HandleFunc("/example/foo/bar", exampleFooBar)
	service.HandleFunc("/example/foo/upload", uploadFile)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func exampleCall(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.Form.Get("name")
	if len(name) == 0 {
		http.Error(w, errors.BadRequest(SERVER_ID, "non content").Error(), 400)
		return
	}

	b, _ := json.Marshal(map[string]interface{}{
		"message": fmt.Sprintf("exampleCall received %s's message", name),
	})

	w.Write(b)
}

func exampleFooBar(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST {
		http.Error(w, errors.BadRequest(SERVER_ID, "require post").Error(), 400)
		return
	}
	if len(r.Header.Get("Content-Type")) == 0 {
		http.Error(w, errors.BadRequest(SERVER_ID, "content type empty").Error(), 400)
		return
	}
	if strings.Trim(r.Header.Get("Content-Type"), " ") != "application/json" {
		http.Error(w, errors.BadRequest(SERVER_ID, "expect appliction/json").Error(), 400)
		return
	}

	bodybytes, _ := ioutil.ReadAll(r.Body)
	var data struct {
		Name string `json:"name"`
	}
	json.Unmarshal(bodybytes, &data)

	b, _ := json.Marshal(map[string]interface{}{
		"message": fmt.Sprintf("exampleFooBar received %s's message", data.Name),
	})

	w.Write(b)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == GET {
		t, _ := template.New("foo").Parse(`
<html>
    <head>
       <title>Upload file</title>
    </head>
    <body>
        <form enctype="multipart/form-data" action="http://localhost:8080/example/foo/upload" method="post">
        <input type="file" name="uploadfile" />
            <br />
                保存目录： <input type="text" name="path" /> 如 /Users/me/Downloads/test/
            <br />
        <input type="submit" name='上传' value="upload" />
        </form>
    </body>
</html>
`)
		t.Execute(w, nil)
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

	path := r.PostForm.Get("path")
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
