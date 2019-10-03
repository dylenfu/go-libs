package ring

import (
	"fmt"
	"testing"
)

var (
	ErrRingEmpty = fmt.Errorf("empty ring")
	ErrRingFull = fmt.Errorf("ring full")
)

type KK struct {
	m string
	n int
}

type Student struct {
	Name string
	Age int
	k KK
}

type Proto struct {
	Name string
	Age int
}

// Ring ring proto buffer.
type Ring struct {
	// read
	rp   uint64
	num  uint64
	mask uint64
	// TODO split cacheline, many cpu cache line size is 64
	// pad [40]byte
	// write
	wp   uint64
	data []Proto
	s Student
}

// NewRing new a ring buffer.
func NewRing(num int) *Ring {
	r := new(Ring)
	r.init(uint64(num))
	return r
}

// Init init ring.
func (r *Ring) Init(num int) {
	r.init(uint64(num))
}

func (r *Ring) init(num uint64) {
	// 2^N
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num &= (num - 1)
		}
		num = num << 1
	}
	r.data = make([]Proto, num)
	r.num = num
	r.mask = r.num - 1
}

// Get get a item from ring.
func (r *Ring) Get() (p *Proto, err error) {
	if r.rp == r.wp {
		return nil, ErrRingEmpty
	}
	p = &r.data[r.rp&r.mask]
	return
}

// GetAdv incr read index.
func (r *Ring) GetAdv() {
	r.rp++
}

// Set get a proto to write.
func (r *Ring) Set() (p *Proto, err error) {
	if r.wp-r.rp >= r.num {
		return nil, ErrRingFull
	}
	p = &r.data[r.wp&r.mask]
	return
}

// SetAdv incr write index.
func (r *Ring) SetAdv() {
	r.wp++
}

// Reset reset ring.
func (r *Ring) Reset() {
	r.rp = 0
	r.wp = 0
	// prevent pad compiler optimization
	// r.pad = [40]byte{}
}

type Channel struct {
	CliProto Ring
}

func TestNewChannel(t *testing.T) {
	ch := new(Channel)
	ch.CliProto.Init(31)
	p , err := ch.CliProto.Set()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
	x := Ring{}
	x.Init(31)
	t.Log(ch)
}