package sourcecode

import (
	"reflect"
	"testing"
	"unsafe"
)

// go test -v github.com/dylenfu/go-libs/sourcecode -run TestSliceHeader
func TestSliceHeader(t *testing.T) {
	data := "hahahajfsljdkfalkdfjalu2ekjflaisjdfakljf29ifnalkgdj"
	ptr1 := unsafe.Pointer(&data)
	ptr2 := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	ptr3 := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)))
	bs := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	t.Log(ptr1, ptr2, ptr3, bs, []byte(data))
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	t.Log(content)
}
