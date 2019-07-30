package skiplist

import "testing"

func TestSkipList(t *testing.T) {
	list := New()
	for i:=0;i<100;i++ {
		list.Insert(i, nil)
	}
	list.PrintList()
}
