package http

func Route(sub string) {
	switch sub {
	case "simple-server":
		SimpleHttpServer()
	}
}
