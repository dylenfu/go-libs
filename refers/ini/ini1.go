package ini

var (
	Num1 int
	Num2 int
)

func init() {
	Num1 = 1
	Num2 = 2
	// println("get into init1")
}

// 所有的init函数在主goroutine(main入口)中执行,通过下划线_引用包，则同一个包的所有init函数都会按顺序执行
func ExecAllInit() {
	println("num1:", Num1)
	println("num2:", Num2)
}
