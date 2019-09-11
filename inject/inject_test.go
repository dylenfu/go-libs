package inject

import (
	"fmt"
	"github.com/facebookgo/inject"
	"testing"
)

type Base struct {
	Id   string
	Desc string
}

type RestApi struct {
	InternalData *Base `inject:"base"`
}

func (s *RestApi) Ring() {
	fmt.Println("inject\t-", "restapi ring", s.InternalData.Id+" "+s.InternalData.Desc)
}

type JsonApi struct {
	InternalData *Base `inject:"base"`
}

func (s *JsonApi) Loop() {
	fmt.Println("inject\t-", "jsonapi loop", s.InternalData.Id+" "+s.InternalData.Desc)
}

type App struct {
	Rest  *RestApi `inject:""`
	Json  *JsonApi `inject:""`
	Other string
}

func (a *App) Sing() {
	a.Rest.Ring()
	a.Json.Loop()
	fmt.Println("inject\t-", "other uninject object", a.Other)
}

// 这里需要注意几点：
// 1.provide的第一个元素为注入对象
// 2.provide的后续元素必须是实例化对象，如果注入interface，那么他应该被实例化后注入
// 3.注入之后Populate会完成Rest&Json的实例化,不需要另外实例化
// 4.被注入结构应该是可访问的(大写)
// 5.如果App中Json没有`inject:""`,则需要实例化app.Json = &datas.JsonApi{base}
func TestRewriteFaceBookInjectDemo(t *testing.T) {
	var graph = inject.Graph{}
	var app App

	base := &Base{"0x12", "base data struct"}
	err := graph.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: base, Name: "base"},
	)
	if err != nil {
		t.Fatal("inject\t-", "inject graph provide error:", err.Error())
	}

	if err := graph.Populate(); err != nil {
		t.Fatal("inject\t-", "inject graph populate error:", err.Error())
	}

	app.Other = "other test"

	app.Sing()
}

// 不能注入interface
func TestInjectInterface(t *testing.T) {
	srv := &AnswerService{}
	stu := &Student{"dylenfu", 30}

	var graph inject.Graph
	graph.Provide(
		&inject.Object{Value: srv},
		&inject.Object{Value: stu, Name: "stu"},
	)
	if err := graph.Populate(); err != nil {
		panic(err)
	}
	srv.Ans.Answer()
}

type AnswerService struct {
	Ans Answerable `inject:"stu"`
}

type Answerable interface {
	Answer()
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) Answer() {
	fmt.Println("inject\t-", "answerable", s.Name, s.Age)
}
