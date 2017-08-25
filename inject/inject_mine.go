package inject

import (
	"github.com/dylenfu/go-libs/inject/datas"
	"github.com/facebookgo/inject"
	"github.com/dylenfu/go-libs/inject/basedatas"
	"log"
)

// 这里需要注意几点：
// 1.provide的第一个元素为注入对象
// 2.provide的后续元素必须是实例化对象，如果注入interface，那么他应该被实例化后注入
// 3.注入之后Populate会完成Rest&Json的实例化,不需要另外实例化
// 4.被注入结构应该是可访问的(大写)
// 5.如果App中Json没有`inject:""`,则需要实例化app.Json = &datas.JsonApi{base}
func RewriteFaceBookInjectDemo() {
	var graph = inject.Graph{}
	var app App

	base := &basedatas.Base{"0x12", "base data strucgt"}
	err := graph.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: base},
	)
	if err != nil {
		log.Panic("inject\t-", "inject graph provide error:", err.Error())
	}

	if err := graph.Populate(); err != nil {
		log.Panic("inject\t-", "inject graph populate error:", err.Error())
	}

	app.Sing()
}

type App struct {
	Rest *datas.RestApi `inject:""`
	Json *datas.JsonApi `inject:""`
}

func (a *App) Sing() {
	a.Rest.Ring()
	a.Json.Loop()
}