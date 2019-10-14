该示例展示了基于元数据的服务发现

使用micro api，而不指定--handler=api,那么handler方式就是rpc。
所以Meta API其实是在RPC模式的基础上，通过在接口层声明端点元数据而指定服务的。

_ = proto.RegisterExampleHandler(service.Server(), new(Example), api.WithEndpoint(&api.Endpoint{
		Name: "Example.Call",
		Description: "",
		Path: []string{"/example"},
		Method: []string{"POST"},
		Handler: rpc.Handler,
	}))

curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example"
{"message":"Meta received john's message"}

curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/foo/bar
2019/10/14 15:30:57 Meta Foo.Bar received request

上述endpoint中，当我们在path内添加"/test", 访问时使用localhost:8080/test,也能mapping到Example.Call。
但是如果将Name稍作修改，Example.Call1,会报错。



