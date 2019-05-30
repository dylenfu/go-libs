package base

import (
	"testing"
	"time"
)

func TestBreakLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
	next1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				t.Log(i, j)
				break next1
			}
		}
	}
}

func TestBreakForLoop(t *testing.T) {
	i := 0
	for {
		if simple(i) {
			i++
			t.Log("========11111")
			time.Sleep(1 * time.Second)
			continue
		}

		for j := 0; j < 10; j++ {
			if j < 5 {
				continue
			} else {
				time.Sleep(1 * time.Second)
				for k := 0; k < 10; k++ {
					if k < 5 {
						continue
					}

					t.Log("========22222")
				}
			}
		}

		t.Log("=========33333")

	}
}

func simple(i int) bool {
	if i > 10 {
		return false
	} else {
		return true
	}
}

func TestMaploop(t *testing.T) {
	m := map[string]int{"1": 11, "2": 22, "3": 33}
	for k, v := range m {
		t.Log(k, v)
	}
}
