package toml

func Route(sub string) {
	switch sub {
	case "quick-start":
		QuickStart()
		break

	case "simple-unmarshal":
		SimpleUnmarshal()
		break
	}
}