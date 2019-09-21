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
	s := []int{2, 3, 4, 5, 6} // len:5 cap:5
	x := s[:2]                //len:2 , cap:5
	y := make([]int, 2, 10)   // len:2, cap:10
	copy(y, x)
	x = y
	x = x[:len(x)+1]
	t.Log(x)

	mm := make([]int, 0, 3)
	mm[0] = 1
	t.Log(mm)
}
