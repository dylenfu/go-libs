package sourcecode

import (
	"testing"
	"unsafe"
)

// go test -v github.com/dylenfu/go-libs/sourcecode -run TestMapSize
// 无论是什么类型的map，使用unsafe.Sizeof获取其大小都是8bytes，这是因为makemap返回的是一个指针
// 查看runtime/internal/sys下PtrSize会发现64位系统下，ptr的大小就是8字节
func TestMapSize(t *testing.T) {
	x := make(map[int]int)
	t.Log(unsafe.Sizeof(x))

	x1 := make(map[string]int)
	t.Log(unsafe.Sizeof(x1))

	x2 := make(map[uint16]uint16)
	t.Log(unsafe.Sizeof(x2))
}
