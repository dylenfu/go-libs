package base

import "testing"

func TestListPtr(t *testing.T) {
	fill1 := Fill{Name: "tom", Heigth: 184}
	fill2 := Fill{Name: "jesse", Heigth: 172}
	list := []Fill{fill1, fill2}
	setFillList(list)

	for _, v := range list {
		println(v.Heigth)
	}
}

type Fill struct {
	Name   string
	Heigth int
}

func setFillList(list []Fill) {
	for _, v := range list {
		v.Heigth = 1
	}
}

// go test -v github.com/dylenfu/go-libs/base -run TestIntPtr
func TestIntPtr(t *testing.T) {
	var x int = 0
	for {
		if x += 1; x > 5 {
			break
		}
		t.Logf("%p %d", &x, x)
	}
}
