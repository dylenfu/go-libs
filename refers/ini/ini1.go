package ini


var (
	Num1 int
	Num2 int
)

func init() {
	Num1 = 1
	Num2 = 2
	println("get into init1")
}

func ExecAllInit() {
	println("num1:", Num1)
	println("num2:", Num2)
}