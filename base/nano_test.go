package base

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// go test -v github.com/dylenfu/algorithm/base -run TestPase
type student struct {
	Name string
	Age  int
}

// 请找到下段代码的问题
// 该测试程序原本意为将stus数组内容赋值到map中
// 但是for range stus中将 stu指针赋值给map的时候有问题
// stu是一个局部变量 其地址不变 只不过在range的过程中 每次将数组的内容逐一复制到该地址而已
func TestPase(t *testing.T) {
	m1 := make(map[string]*student)
	m2 := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for k, stu := range stus {
		m1[stu.Name] = &stu
		m2[stu.Name] = &stus[k]
		fmt.Println(stu)
	}
	for k, v := range m1 {
		fmt.Printf("student :%s, name:%s, age:%d\r\n", k, v.Name, v.Age)
	}
	for k, v := range m2 {
		fmt.Printf("student :%s, name:%s, age:%d\r\n", k, v.Name, v.Age)
	}
}

func TestRuntime(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	//wg.Add(20)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 测试结构体继承
// 打印:
// people showA
// people showB
// teacher showB
type People struct{}

func (p *People) ShowA() {
	fmt.Println("people showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("people showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestShow(t *testing.T) {
	teacher := Teacher{}
	teacher.ShowA()
	teacher.ShowB()
}

// 测试 下面的代码会显示异常吗
// 会
func TestSelect(t *testing.T) {
	runtime.GOMAXPROCS(1)

	string_chan := make(chan string, 1)
	int_chan := make(chan int, 1)

	string_chan <- "say hello"
	int_chan <- 1

	select {
	case value := <-string_chan:
		fmt.Println(value)
	case value := <-int_chan:
		fmt.Println(value)
		// default:
	}
}
