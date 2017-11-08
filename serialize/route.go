package serialize

import "github.com/dylenfu/go-libs/serialize/protodata"

func Route(sub string) {
	switch sub {
	case "simple-proto-2":
		protodata.SimpleProto2()

	case "simple-proto-3":
		protodata.SimpleProto3()
	}
}
