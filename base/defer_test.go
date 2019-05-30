package base

import (
	"testing"
	"fmt"
)

// defer的实现原理，在runtime2.go文件中定义了_defer的数据结果，
// 与func类似,拥有程序计数器，栈指针，函数地址指针，此外，有一个指向自身的defer指针
// 任意一个gorutine中，defer都是以链表的形式进行链接的
// 在panic.go中，有两个函数deferproc&deferretrun
// 在deferproc中用来声明defer，将defer对应的func插入到链表头部，
// 在deferreturn中，从链表头部取出一个defer函数进行执行
//
// defer的执行在panic之前
// defer对应的函数入参不变

// go test -v github.com/dylenfu/go-libs/base -run TestDefer1
// @result 3 2 1 & panic content
func TestDefer1(t *testing.T) {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	panic("panic now")
}

// 打印结果1
// defer执行的操作参数a,定义在a = 2之前
func TestDefer2(t *testing.T) {
	a := 1
	defer fmt.Println(a)
	a = 2
}

// defer的入参为数组对象指针
func TestDefer3(t *testing.T) {
	list := [3]int{7, 9, 3}
	defer printArray(&list)
	list[2] = 5
}

func printArray(list *[3]int) {
	for _, v := range list {
		fmt.Println(v)
	}
}

//
// 下面几个例子讲defer对return的影响
// return在go里的操作并非原子化的 他包含两个步骤：
// 1.将结果存入栈中
// 2.跳转
// defer的执行在跳转之前
//
// TestDefer4的打印结果为1 TestDefer5的打印结果为0
// 因为deferFuncReturn1是具名返回值，最后的返回操作return相当于:
// i = 0; i++; return i;
// 而deferFuncReturn2是匿名返回值，最后的返回操作相当于:
// result := i; i++; return result;
//
// 同理:
// TestDefer6的打印结果为0 TestDefer7的打印结果为1
// deferFuncReturn3的return操作相当于:
// result = 0; i++; return result; 字面量直接存入栈中
// deferFuncReturn4的return操作相当于:
// i = 0; i++; return i;
func TestDefer4(t *testing.T) {
	fmt.Println(deferFuncReturn1())
}

func TestDefer5(t *testing.T) {
	fmt.Println(deferFuncReturn2())
}

func TestDefer6(t *testing.T) {
	fmt.Println(deferFuncReturn3())
}

func TestDefer7(t *testing.T) {
	fmt.Println(deferFuncReturn4())
}

// 具名返回值 i int
func deferFuncReturn1() (i int) {
	i = 0
	defer func() {
		i++
	}()
	return i
}

// 匿名返回值
func deferFuncReturn2() int {
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
