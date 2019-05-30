package base

import (
	"log"
	"testing"
)

func TestInterface(t *testing.T) {
	r := &Rectangle{3.15, 3.0}
	use(r)
}

// implement interface Geo
type Geo interface {
	Area() float32
	Perimeter() float32
}

func use(g Geo) {
	log.Println("interface\t-", "geo area:", g.Area())
	log.Println("interface\t-", "geo perimeter", g.Perimeter())
}

type Rectangle struct {
	L float32
	W float32
}

func (r *Rectangle) Area() float32 {
	return r.L * r.W
}

func (r *Rectangle) Perimeter() float32 {
	return 2*r.W + 2*r.L
}
