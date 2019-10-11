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
