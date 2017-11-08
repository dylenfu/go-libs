package mysql

func Route(sub string) {
	switch sub {
	case "simple-orm":
		SimpleOrm()
	}
}
