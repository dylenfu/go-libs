package lintcode

import "testing"

func TestDot(t *testing.T) {
	a := []int{}
	b := []int{}
	t.Log(dotProduct(a, b))
}

func dotProduct (A []int, B []int) int {
	// Write your code here
	lenA := len(A)
	lenB := len(B)
	if lenA <= 0 || lenB <= 0 || lenA != lenB {
		println("there is not dot product")
		return -1
	}

	ret := 0
	for i := 0; i < lenA; i++ {
		ret += A[i] * B[i]
	}
	return ret
}
