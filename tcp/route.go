package tcp


func Route(sub string) {
	switch sub {
	case "client":
		client()

	case "server":
		server()
	}
}