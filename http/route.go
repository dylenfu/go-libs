package http

import (
	"github.com/dylenfu/go-libs/http/simple"
	"github.com/dylenfu/go-libs/http/jsonrest"
	"github.com/dylenfu/go-libs/http/ref"
)

func Route(sub string) {
	switch sub {
	case "simple-server":
		simple.SimpleHttpServer()

	case "simple-reflect":
		ref.SimpleReflect()

	case "jsonrest-hello":
		jsonrest.Hello()

	case "jsonrest-simple-route":
		jsonrest.SimpleRoute()

	case "jsonrest-simple-reflect":
		jsonrest.Reflect()
	}
}
