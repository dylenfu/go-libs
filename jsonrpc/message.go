package jsonrpc

type ReqMessage struct {
	Id int `json:id`
	Name string `json:name`
}

type RespMessage struct {
	Ok bool `json:ok`
	Id int `json:id`
	Msg string `json:msg`
}