package base

import (
	"fmt"
	"testing"
)

// defer的执行顺序问题：
// defer的实现原理，在runtime2.go文件中定义了_defer的数据结果，
// 与func类似,拥有程序计数器，栈指针，函数地址指针，此外，有一个指向自身的defer指针
// 任意一个gorutine中，defer都是以链表的形式进行链接的
// 在panic.go中，有两个函数deferproc&deferretrun
// 在deferproc中用来声明defer，将defer对应的func插入到链表头部，
// 在deferreturn中，从链表头部取出一个defer函数进行执行
// 当遇到panic，panic代码后的defer不会执行
// go test -v github.com/dylenfu/go-libs/base -run TestDefer1
// @result 3 2 1
func TestDeferLink(t *testing.T) {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func TestDeferPanic(t *testing.T) {
	defer fmt.Println(1)
	panic("panic now")
	defer fmt.Println(2)
}

func TestDeferReturned(t *testing.T) {
	t.Log(deferAfterReturn())
}

func deferAfterReturn() (x int) {
	x = 5
	return
	defer func() { x = 3 }()
	return
}

// defer执行时函数入口参数的问题
// defer对应的函数入参不变
// 1.当传入的参数是常量时，defer使用常量；如果是普通变量，defer使用的是变量对应的常量值
//   如TestDeferVar defer执行的操作参数a,定义在a = 2之前
// 2.当defer使用的是指针时，参考1可知，指针不变的情况下，defer会使用后续代码中变更的引用内容
//   如TestDeferArrayPtr，打印结果是&[7, 9, 5]
// 3.当defer遇到匿名函数，受用func(){}()将代码块包起来了，情况与2同，
//   如TestDeferStructPtr，当m是结构体对象时，打印结果中age是30，当m是结构体对象引用时，age是40
func TestDeferVar(t *testing.T) {
	a := 1
	defer fmt.Println(a)
	a = 2
}

// defer的入参为数组对象指针
func TestDeferArrayPtr(t *testing.T) {
	list := [3]int{7, 9, 3}
	defer fmt.Println(&list)
	list[2] = 5
}

func TestDeferStructPtr(t *testing.T) {
	type T struct {
		name string
		age  int
	}

	m := T{
		name: "dylenfu",
		age:  30,
	}

	defer func(d T) {
		fmt.Printf("name is %s, age is %d:", d.name, d.age)
	}(m)

	m.age = 40
}

//
// 下面几个例子讲defer对return的影响
// return在go里的操作并非原子化的 他包含两个步骤：
// 1.将结果存入栈中
// 2.跳转
// defer的执行在跳转之前
//
// TestDeferNamedReturn的打印结果为1
// TestDeferAnonymosReturn的打印结果为0
// 因为TestDeferNamedReturn是具名返回值，最后的返回操作return相当于:
// i = 0; i++; return i;
// 而TestDeferAnonymosReturn是匿名返回值，最后的返回操作相当于:
// result := i; i++; return result;
//
// 同理:
// TestDefer6的打印结果为0 TestDefer7的打印结果为1
// deferFuncReturn3的return操作相当于:
// result = 0; i++; return result; 字面量直接存入栈中
// deferFuncReturn4的return操作相当于:
// i = 0; i++; return i;
func TestDeferNamedReturn(t *testing.T) {
	fmt.Println(deferNamedReturn())
}

func TestDeferAnonymosReturn(t *testing.T) {
	fmt.Println(deferAnonymosReturn())
}

func TestDefer6(t *testing.T) {
	fmt.Println(deferFuncReturn3())
}

func TestDefer7(t *testing.T) {
	fmt.Println(deferFuncReturn4())
}

// 具名返回值 i int
func deferNamedReturn() (i int) {
	i = 0
	defer func() {
		i++
	}()
	return i
}

// 匿名返回值
func deferAnonymosReturn() int {
	i := 0
	defer func() {
		i++
	}()
	return i
}

func deferFuncReturn3() int {
	var i int
	defer func() {
		i++
	}()
	return 0
}

func deferFuncReturn4() (i int) {
	defer func() {
		i++
	}()
	return 0
}

//
// 下面的两个例子用来证明defer 调用匿名函数&具名函数的区别
// defer调用具名函数时 函数作为参数在外围先执行
// 在调用匿名函数时 所有内容都要在退出前执行
func TestDeferNamedFunc(t *testing.T) {
	defer deferOutFunc(deferInnerFunc())
	fmt.Println("starting print")
}

func TestDeferAnonymosFunc(t *testing.T) {
	defer func() {
		deferOutFunc(deferInnerFunc())
	}()
	fmt.Println("starting print")
}

func deferInnerFunc() string {
	fmt.Println("inner func")
	return "abc"
}

func deferOutFunc(str string) {
	fmt.Printf("out func:%s", str)
}
