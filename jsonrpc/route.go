package jsonrpc

func Route(sub string) {
	switch sub {
	case "simple-server":
		NewServer1()
		break

	case "aync-call":
		AyncCall()
		break

	case "sync-call":
		SyncCall()
		break
	}
}
