package sourcecode

import (
	"testing"
	"unsafe"
	"reflect"
)

func TestSlice(t *testing.T) {
	s := make([]int, 10)
	for i := 0; i < 10; i++ {
		s[i] = i + 1000
	}

	s = s[:5]
	t.Log("kk")
}

func TestSlice2(t *testing.T) {
	s := []int{2, 3, 4, 5, 6} // len:5 cap:5
	x := s[:2]                //len:2 , cap:5
	y := make([]int, 2, 10)   // len:2, cap:10
	copy(y, x)
	x = y
	x = x[:len(x)+1]
	t.Log(x)

	mm := make([]int, 0, 3)
	mm[0] = 1
	t.Log(mm)
}

// go test -v github.com/dylenfu/go-libs/sourcecode -run TestArrayModifyWithUnsafePointerOffset
// desc: 测试数组通过unsafe.Sizeof元素长度作为偏移量，查询或者修改对应位置的数据内容
// 64位机器上的打印结果是:
// slice_test.go:37: 24
// slice_test.go:39: 8
// slice_test.go:41: 2
// slice_test.go:44: 10
// slice_test.go:53: 1316808
// 24，8分别是整个array的长度，每个int类型的长度8，3个元素的总长度是24，说明数组只存储了3个int类型的数据，没有别的内容
// 2是数组下标为1的内容，通过unsafe获取指定偏移量对应的元素内容
// 10是修改该地址位置数据内容的结果
// 最后数组越界打印的数据，这里已经访问到内存其他的位置了，不属于arr的范畴
func TestArrayModifyWithSizeof(t *testing.T) {
	var arr = [3]int{1, 2, 3}

	t.Log(unsafe.Sizeof(arr))
	size := unsafe.Sizeof(arr[0])
	t.Log(size)

	t.Log(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + 1*size)))
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + 1*size)) = 10

	t.Log(arr[1])
	t.Log(*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + 8 * size)))
}

// desc: 测试数组拷贝, 打印结果如下
// slice_test.go:60: 0xc0000be448
// slice_test.go:63: 0xc0000be480
// 可以看出，数组传递后，元素内容存储地址不一样，这说明array的传递是值传递，而不是指针传递
func TestArrayCopy(t *testing.T) {
	var arr = [1]int{5}
	t.Log(&arr[0])

	arr1 := arr
	t.Log(&arr1[0])
}

// desc: 测试slice属性(len, cap)在内存中的位置
// slice_test.go:77: 824634442784 824634459232
// slice_test.go:78: 100
// slice_test.go:79: 824634442792 1
// slice_test.go:80: 824634442800 2
func TestSliceAttributeInMemoryIdx(t *testing.T) {
	s := make([]int, 1, 2)
	s[0] = 100

	size := unsafe.Sizeof(int(0))
	t.Log(uintptr(unsafe.Pointer(&s)), *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)))))
	t.Log(*(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s)))))
	t.Log(uintptr(unsafe.Pointer(&s)) + size, *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size)))
	t.Log(uintptr(unsafe.Pointer(&s)) + size * 2, *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size * 2)))
}

// desc: 测试slice的值拷贝到array
// slice_test.go:91: [0 0 20 0 0 0 0 0 0 100 0 0 0 0 0 0 0 0 0 0]
func TestSliceCapMatchArrayLen(t *testing.T) {
	s := make([]int, 10, 20)
	s[2] = 20
	s[9] = 100

	t.Log(*(*[20]int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&s)))))
}

// desc: 测试bool slice grow
// 观察到新的s的capacity是16，既不是8，也不是12
func TestBoolSliceGrow(t *testing.T) {
	s := make([]bool, 4, 4)
	s1 := make([]bool, 8, 8)
	s = append(s, s1...)
}

// desc: 测试slice拷贝
// slice_test.go:114: 824634867744 824634867776    // s,s1地址
// slice_test.go:115: 824634867752 7               // s.Len 地址以及地址内容
// slice_test.go:116: 824634867784 2               // s1.Len 地址以及地址内容
// slice_test.go:117: 824634900552 824634900552    // s[1],s1[1]元素地址
// slice_test.go:119: 7                            // s.cap
// slice_test.go:121: 18                           // 扩容后s.cap
// slice_test.go:122: 824635097096                 // 扩容后s[1]元素地址
func TestSliceCopy(t *testing.T) {
	s := make([]int, 7, 7)
	s1 := s[:2]

	size := unsafe.Sizeof(0)
	t.Log(uintptr(unsafe.Pointer(&s)), uintptr(unsafe.Pointer(&s1)))
	t.Log(uintptr(unsafe.Pointer(&s)) + size, *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + size)))
	t.Log(uintptr(unsafe.Pointer(&s1)) + size, *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s1)) + size)))
	t.Log(uintptr(unsafe.Pointer(&s[1])), uintptr(unsafe.Pointer(&s1[1])))

	t.Log(cap(s))
	s = append(s, make([]int, 10, 10)...)
	t.Log(cap(s))
	t.Log(uintptr(unsafe.Pointer(&s[1])))
}

// desc: 测试nil slice
// s0只声明，他是一个nil slice，其他的都不是，make后会给定一个地址，
// slice_test.go:139: {0 0 0} [] true
// slice_test.go:140: {23829736 0 0} [] false
// slice_test.go:141: {23829736 0 0} [] false
// slice_test.go:142: {824635105280 0 100} [] false
// slice_test.go:143: {824633835920 0 10} [] false
func TestNilSlice(t *testing.T) {
	var s0 []int
    s1 := make([]int, 0)
    s2 := make([]int, 0, 0)
    s3 := make([]int, 0, 100)

    arr := [10]int{}
    s4 := arr[:0]

	t.Log("s0", *(*reflect.SliceHeader)(unsafe.Pointer(&s0)), s0, s0 == nil)
    t.Log("s1", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), s1, s1 == nil)
    t.Log("s2", *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), s2, s2 == nil)
    t.Log("s3", *(*reflect.SliceHeader)(unsafe.Pointer(&s3)), s3, s3 == nil)
    t.Log("s4", *(*reflect.SliceHeader)(unsafe.Pointer(&s4)), s4, s4 == nil)
}

func TestSlicePost(t *testing.T) {
	s := make([]int, 1, 2)
	st(s, t)
	t.Log(s)
}

func st(s []int, t *testing.T) {
	s[0] = 10
}