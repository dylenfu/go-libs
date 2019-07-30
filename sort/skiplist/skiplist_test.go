package skiplist

import "testing"

func TestSkipList(t *testing.T) {
	list := New()
	for i:=0;i<2;i++ {
		list.Insert(i + 1000, nil)
	}
	list.PrintList()
}
