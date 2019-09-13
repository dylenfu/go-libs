package base

import (
	"testing"
	"math"
)

// go test -v github.com/dylenfu/go-libs/base -run TestAddWithBinararyAction
func TestAddWithBinararyAction(t *testing.T) {
	if (3 + 4) != (3 | 4) {
		t.Fatal("(3 + 4) != (3 | 4)")
	}

	if math.MaxInt32 != (1 << 31 - 1) {
		t.Log("math.MaxInt32 != (1 << 32 - 1)")
	}
	if math.MinInt32 != -(1 << 31) {
		t.Log("math.MinInt32 != -(1 << 31)")
	}

	if (3 ^ 4) != (3 + 4) {
		t.Log("(3 ^ 4) != (3 + 4)")
	}

	if aplusb(100, -100) != 0 {
		t.Fatal("aplusb(100, -100) != 0")
	}
}

// 使用二进制方式实现a + b
// a ^ b 如果在计算过程中没有进位 那么 a + b == a ^ b
// a & b << 1 但是如果有进位，那么对进位左移一位，同时非进位置位为0就是进位的值
func aplusb(a, b int) int {
	if b == 0 {
		return a
	}

	return aplusb((a ^ b), (a & b << 1))
}
