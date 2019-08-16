package min

type Elem struct {
	Priority int
	Value    interface{}
	Index    int
}

// todo some bug in Pop&Push
type Heap []*Elem

func (h *Heap) Push(e *Elem) {
	h.push(e)
	n := h.Len() - 1
	h.up(n)
}

func (h *Heap) Pop() *Elem {
	n := h.Len() - 1
	h.Swap(0, n)
	h.down(0, n)
	return h.pop()
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h Heap) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Remove(i int) interface{} {
	n := len(h) - 1
	if n != i {
		h.Swap(i, n)
		h.down(i, n)
		h.up(i)
	}
	return h.Pop()
}

func (h *Heap) push(e *Elem) {
	n := h.Len()
	e.Index = n
	*h = append(*h, e)
}

func (h *Heap) pop() *Elem {
	old := *h
	n := old.Len()
	e := old[n-1]
	e.Index = -1 // for safety
	*h = old[0 : n-1]
	return e
}

func (h *Heap) up(i0 int) bool {
	i := i0
	for {
		j := (i - 1) / 2             // parant
		if i == j || !h.Less(i, j) { // i和j都是0或者i不在j前面
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i < i0
}

func (h *Heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := (i * 2) + 1
		if j1 >= n || j1 < 0 { // j1 must < n, besides if i is big enough, j may be overflow;
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && !h.Less(j2, j1) {
			j = j2 // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}

	return i > i0
}
