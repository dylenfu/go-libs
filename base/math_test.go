package base

import (
	"math"
	"math/big"
	"testing"
)

func TestSimpleMath(t *testing.T) {
	t.Log("9开平方根", math.Sqrt(9))
	t.Log("10的3次方", math.Pow(10, 3))
	t.Log("27开3次方根", math.Pow(27, 0.3333333333333))
}

func TestBigRatAndBigInt(t *testing.T) {
	amount, _ := new(big.Int).SetString("100000000000000000000000000001", 10)
	t.Log(amount)

	amountS := big.NewRat(3, 2) // 1.5
	amountB := new(big.Rat).SetFrac(amount, big.NewInt(1))

	m := new(big.Rat).Mul(amountB, amountS)
	r := m.Mul(m, big.NewRat(1, 1))
	t.Log(r.Num())
}

func TestUseE(t *testing.T) {
	a := 1.0320388897123451e18
	t.Log(a)
}

func TestBigRatToFloat(t *testing.T) {
	a := big.NewRat(1, 3)
	t.Log(a.Float64())
}
