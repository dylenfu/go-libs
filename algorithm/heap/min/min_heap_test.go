package min

import (
	"fmt"
	"testing"
)

func TestHeap_Push(t *testing.T) {
	h := &Heap{}
	for i := 10; i > 0; i-- {
		h.Push(&Elem{i, nil, 0})
	}
	for i := 0; i < h.Len(); i++ {
		e := h.Pop()
		t.Log(fmt.Sprintf("elem [index %d] [priority %d]", e.Index, e.Priority))
	}
}
