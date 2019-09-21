package pqueue

// TODO(fukun): add mutex???
type PQueue struct {
	list []Entry
}

const (
	minSize            = 1 << 4
	maxSize            = 1 << 10
	shrinkThreshold    = 1 << 5
	overflowThreshold  = 1 << 5 // maxSize - currentSize
	overflowRemoveSize = 1 << 5 // maxSize - currentSize
)

func adjustSize(size int) int {
	if size < minSize {
		return minSize
	} else if size > maxSize {
		return maxSize
	} else {
		return nextPowerOf2(size)
	}
}

func NewPriorityQueue(capacity int) *PQueue {
	capacity = adjustSize(capacity)
	return &PQueue{
		list: make([]Entry, 0, capacity),
	}
}

func (q *PQueue) Len() int {
	return len(q.list)
}

func (q *PQueue) Cap() int {
	return cap(q.list)
}

func (q *PQueue) Less(i, j int) bool {
	return q.list[i].Less(q.list[j])
}

func (q *PQueue) Swap(i, j int) {
	q.list[i], q.list[j] = q.list[j], q.list[i]
}

// Push add x as element Len(n)
func (q *PQueue) Push(e Entry) {
	n := q.Len()
	c := q.Cap()

	reduced := q.overflow(n, c)
	q.expand(n, c)

	if reduced {
		n = q.Len()
	}

	q.list = q.list[:n+1]
	q.list[n] = e
	q.up(n)
}

// Pop remove and return element as 0
func (q *PQueue) Pop() Entry {
	n := len(q.list)
	c := cap(q.list)

	q.Swap(0, n-1)
	q.down(0, n-1)

	entry := q.list[n-1]
	q.list = q.list[:n-1]

	q.shrink(n, c)

	return entry
}

// remove remove and return element as Len() - 1
func (q *PQueue) remove() Entry {
	n := len(q.list)
	c := cap(q.list)

	entry := q.list[n-1]
	q.list = q.list[:n-1]

	q.shrink(n, c)

	return entry
}

// expand multiply and expand capacity while length near 2^N
func (q *PQueue) expand(n, c int) bool {
	if n+2 < c {
		return false
	}

	npq := make([]Entry, n, c*2)
	copy(npq, q.list)
	q.list = npq

	return true
}

// overflow pop some elements from queue
func (q *PQueue) overflow(n, c int) bool {
	if n+overflowThreshold < maxSize {
		return false
	}

	for i := 0; i < overflowRemoveSize; i++ {
		q.remove()
	}

	return true
}

// shrink reduce the capacity by half while length near 2^(N-1)
// e.g c = 64, n = 13; after shrink c = 32, n = 13
// e.g c = 128, n = 31; after shrink c = 64, n = 31
func (q *PQueue) shrink(n, c int) bool {
	if n < (c/4) && c > shrinkThreshold {
		npq := make([]Entry, n, c/2)
		copy(npq, q.list)
		q.list = npq
		return true
	}

	return false
}

// up x floating up and swap with parent until x >= parent.
// initial index j0 is bigger if floating happen actually
func (q *PQueue) up(j0 int) bool {
	j := j0
	for {
		i := (j - 1) / 2 // parent element
		if i == j || !q.Less(j, i) {
			break
		}
		q.Swap(i, j)
		j = i
	}
	return j < j0
}

// down x sink down and swap with child until x <= child
// n is length of queue, initial index i0 is smaller
func (q *PQueue) down(i0, n int) bool {
	i := i0
	for {
		j1 := i*2 + 1 // left child
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j + 1; j2 < n && q.Less(j2, j1) {
			j = j2
		}
		if !q.Less(j, i) {
			break
		}
		q.Swap(i, j)
		i = j
	}

	return i > i0
}

func nextPowerOf2(n int) int {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}
