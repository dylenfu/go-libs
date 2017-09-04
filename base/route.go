package base

func Route(sub string) {
	switch sub {
	case "hi":
		println("Hello IPFS")
		break

	case "simple-struct":
		t := &Tee{3, "hello"}
		println(t.GetTeeNum())
		println(t.GetTeeName())
		break

	case "simple-channel":
		ChannelDemo()
		break

	case "simple-channel-wait":
		SimpleChannelDemo()
		break

	case "simple-interface":
		InterfaceDemo()

	case "simple-reflect1":
		ReflectDemo1()
		break

	case "simple-reflect2":
		ReflectDemo2()
		break

	case "simple-reflect3":
		ReflectDemo3()
		break

	case "simple-reflect4":
		ReflectDemo4()
		break

	case "simple-reflect5":
		ReflectDemo5()
		break

	case "simple-reflect6":
		ReflectDemo6()
		break

	case "reflect-judge":
		JudgeType()
		break

	case "simple-path":
		SimplePath()
		break

	case "simple-math":
		SimpleMath()
		break

	default:
		break
	}
}
