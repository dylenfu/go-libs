package inject

func Route(sub string) {
	switch sub {
	case "quick-start":
		SimpleInject()
		break

	case "simple-mine":
		RewriteFaceBookInjectDemo()
		break

	case "simple-interface":
		InjectInterface()
		break

	}
}