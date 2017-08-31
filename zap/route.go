package zap

func Route(sub string) {
	switch sub {
	case "simple-quick-start":
		SimpleZapLogger()
		break

	case "simple-save":
		SimpleSavingZapLogger()
		break

	case "simple-multi-save":
		MultipleSavingZapLogger()
		break

	case "simple-logger-print":
		SimpleLoggerAndPrint()
		break
	}
}