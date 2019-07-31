package skiplist

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	list := New()
	for i := 1; i <= 100; i++ {
		//if i == 6 {
		//	fmt.Println("debug here")
		//}
		list.Insert(i, nil)
	}
	list.PrintList()
}
