package protodata

import (
	"github.com/golang/protobuf/proto"
)


/*
在main函数所在入口生成pb.go protoc --go_out=. serialize/protodata/*.proto
在pkg里生成pb.go protoc --go_out=. *.proto
*/

func SimpleProto2() {
	test := &Test{}
	test.Id = proto.Int32(1)
	test.Opt = proto.Int32(2)
	test.Str = proto.String("test")

	println("test.Id", test.GetId())
	println("test.Opt", test.GetOpt())
	println("test.Str", test.GetStr())
}

func SimpleProto3() {
	test := &Test1{}
	test.Ed = Test1_X
	test.Page = 3
	test.Names = []string{"name1", "name2"}
	test.Users = map[string]int32{"name1":32, "name2":33}

	println("test.Ed", test.Ed)
	println("test.Page", test.Page)
	println("test.Names", test.Names)
	println("test.Users", test.Users)
}

