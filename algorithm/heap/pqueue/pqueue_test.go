package pqueue

import "testing"

type Msg struct {
	idx int
}

func (m *Msg) Value() interface{} {
	return m.idx
}

// 最小堆反过来即可
func (m *Msg) Less(x Entry) bool {
	return m.idx > x.(*Msg).idx
}

func iton(m Entry) int {
	switch x := m.Value().(type) {
	case int:
		return x
	default:
		panic("invalid num type")
	}
	return 0
}

func TestSimpleMaxHeap(t *testing.T) {
	q := NewPriorityQueue(8)

	for i := 0; i < 16; i++ {
		q.Push(&Msg{idx: 16 - i - 1})
	}

	if q.Len() != 16 || q.Cap() != 32 {
		t.Fatal("length or cap error")
	}

	for i := 0; i < 16; i++ {
		e := q.Pop()
		if m, ok := e.(*Msg); !ok || m.idx != 16-i-1 {
			t.Fatal("index error")
		}
	}
}

func TestRepeatedElements(t *testing.T) {
	q := NewPriorityQueue(8)

	q.Push(&Msg{idx: 18})
	q.Push(&Msg{idx: 6})
	q.Push(&Msg{idx: 9})
	q.Push(&Msg{idx: 7})
	q.Push(&Msg{idx: 2})
	q.Push(&Msg{idx: 7})
	q.Push(&Msg{idx: 6})
	q.Push(&Msg{idx: 3})

	if iton(q.Pop()) != 18 {
		t.Fatal("iton(q.Pop()) != 18")
	}
	if iton(q.Pop()) != 9 {
		t.Fatal("iton(q.Pop()) != 9")
	}
	if iton(q.Pop()) != 7 {
		t.Fatal("iton(q.Pop()) != 7")
	}
	if iton(q.Pop()) != 7 {
		t.Fatal("iton(q.Pop()) != 7")
	}
	if iton(q.Pop()) != 6 {
		t.Fatal("iton(q.Pop()) != 6")
	}
	if iton(q.Pop()) != 6 {
		t.Fatal("iton(q.Pop()) != 6")
	}
	if iton(q.Pop()) != 3 {
		t.Fatal("iton(q.Pop()) != 3")
	}
	if iton(q.Pop()) != 2 {
		t.Fatal("iton(q.Pop()) != 2")
	}
}

func TestExpand(t *testing.T) {
	q := NewPriorityQueue(13)

	if q.Cap() != 16 {
		t.Fatal("initialize queue error")
	}

	// 1024 + 128
	for i := 0; i < maxSize+128; i++ {
		q.Push(&Msg{i})
		x := i + 1
		switch x {
		case 16:
			if q.Len() != 16 || q.Cap() != 32 {
				t.Fatal("q.Len() != 16 || q.Cap() != 32")
			}
		case 32:
			if q.Len() != 32 || q.Cap() != 64 {
				t.Fatal("q.Len() != 32 || q.Cap() != 64")
			}
		case 64:
			if q.Len() != 64 || q.Cap() != 128 {
				t.Fatal("q.Len() != 64 || q.Cap() != 128")
			}
		case 128:
			if q.Len() != 128 || q.Cap() != 256 {
				t.Fatal("q.Len() != 128 || q.Cap() != 256")
			}
		case 256:
			if q.Len() != 256 || q.Cap() != 512 {
				t.Fatal("q.Len() != 256 || q.Cap() != 512")
			}
		case 512:
			if q.Len() != 512 || q.Cap() != 1024 {
				t.Fatal("q.Len() != 512 || q.Cap() != 1024")
			}
		}
		if x+overflowThreshold == maxSize {
			if q.Len() != maxSize-overflowThreshold || q.Cap() != maxSize {
				t.Fatal("q.Len() != maxSize - overflowRemoveSize || q.Cap() != 1024")
			}
		}
		if x+overflowThreshold == maxSize+1 {
			if q.Len() != maxSize-overflowThreshold-overflowRemoveSize+1 || q.Cap() != maxSize {
				t.Fatal("q.Len() != maxSize - overflowThreshold - overflowRemoveSize || q.Cap() != maxSize")
			}
		}
	}

	if q.Len() != maxSize-overflowThreshold || q.Cap() != 1024 {
		t.Fatal("q.Len() != maxSize - overflowThreshold || q.Cap() != 1024")
	}

	for i := 0; i < q.Len(); i++ {
		x := iton(q.Pop())
		y := maxSize + 128 - 1 - i
		if x == 1109 || y == 1109 {
			t.Log(1)
		}
		if x != y {
			t.Fatalf("iton(q.Pop()) != maxSize + 128 - 1 - i, %d != %d \n", x, y)
		}
	}
}
