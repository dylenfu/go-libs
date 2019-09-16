package base

import "testing"

func TestSlice(t *testing.T) {
	s := make([]int, 10)
	for i := 0; i < 10; i++ {
		s[i] = i + 1000
	}

	s = s[:5]
	t.Log("kk")
}

func TestSlice2(t *testing.T) {
	s := []int{2, 3}
	t.Log(s[:2])
}