package sourcecode

import (
	"testing"
	"unsafe"
)

// desc: 测试不同类型指针通过unsafe.Pointer进行转换
// go的指针不允许跨类型转换, 转换时需要使用unsafe
// go test -v github.com/dylenfu/go-libs/sourcecode -run TestUnsafePointerConverter
func TestUnsafePointerConverter(t *testing.T) {
	i := 10
	ip := &i

	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3.0
	t.Log(fp, &fp, *fp)
}

// desc: 测试unsafe.Offsetof和uintptr，通过偏移量获取结构体字段内容
// go的指针和unsafe.Pointer都不能计算偏移量，这样就没法得到固定地址的内存内容了
// 可以将指针转换为uintptr, 该指针类型可以进行偏移量计算
func TestUnsafePointerOffset(t *testing.T) {
	type user struct {
		name string
		age  int
	}

	u := new(user)
	t.Log(&u, u)

	pname := (*string)(unsafe.Pointer(u))
	*pname = "dylen"

	// 注意: 这里只是为了方便看清楚代码 所以将
	// 一般而言最好是连在一起写，防止临时变量被gc
	if false {
		start := uintptr(unsafe.Pointer(u))
		offset := unsafe.Offsetof(u.age)
		end := start + offset
		page := (*int)(unsafe.Pointer(end))
		*page = 20
		t.Log(start, offset, end, u)
	} else {
		page := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
		*page = 30
		t.Log(u)
	}
}

// 测试结果可以看出，结构体在内存中存储是一片连续的内存，以8字节为单位存储每个字段的指针
//unsafe_pointer_test.go:61: size of s 8
//unsafe_pointer_test.go:69: age offset 824634224096 15
//unsafe_pointer_test.go:70: pc offset 824634224104 3
//unsafe_pointer_test.go:71: class offset 824634224112 6
//unsafe_pointer_test.go:72: grade offset 824634224120 3
//在内存中的总长度为 24+8，即便是uint16，存储的也是8字节指针
func TestSimpleStructPointerAndIndex(t *testing.T) {
	type scase struct {
		age   uint16
		pc    uintptr
		class int64
		grade int64
	}

	s := &scase{age: 15, pc: uintptr(3), class: 6, grade: 3}

	t.Log("size of s", unsafe.Sizeof(s))

	a0 := uintptr(unsafe.Pointer(s))
	a1 := a0 + unsafe.Offsetof(s.age)
	a2 := a0 + unsafe.Offsetof(s.pc)
	a3 := a0 + unsafe.Offsetof(s.class)
	a4 := a0 + unsafe.Offsetof(s.grade)

	t.Log("age offset", a1, *(*uint16)(unsafe.Pointer(a1)))
	t.Log("pc offset", a2, *(*uintptr)(unsafe.Pointer(a2)))
	t.Log("class offset", a3, *(*int64)(unsafe.Pointer(a3)))
	t.Log("grade offset", a4, *(*int64)(unsafe.Pointer(a4)))
}
