package base

func Route(sub string) {
	switch sub {
	case "hi":
		println("Hello IPFS")

	case "simple-string":
		SimpleStringType()

	case "simple-atomic":
		SimpleAtomic()

	case "simple-struct":
		t := &Tee{3, "hello"}
		println(t.GetTeeNum())
		println(t.GetTeeName())

	case "simple-channel":
		ChannelDemo()

	case "simple-channel-wait":
		SimpleChannelDemo()

	case "simple-interface":
		InterfaceDemo()

	case "simple-reflect1":
		ReflectDemo1()

	case "simple-reflect2":
		ReflectDemo2()

	case "simple-reflect3":
		ReflectDemo3()

	case "simple-reflect4":
		ReflectDemo4()

	case "simple-reflect5":
		ReflectDemo5()

	case "simple-reflect6":
		ReflectDemo6()

	case "reflect-judge":
		JudgeType()

	case "simple-path":
		SimplePath()

	case "simple-os-args":
		SimpleOsArgs()
		break

	case "simple-math":
		SimpleMath()
	}
}
