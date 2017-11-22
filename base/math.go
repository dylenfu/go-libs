package base

import (
	"log"
	"math"
	"math/big"
)

func SimpleMath() {
	log.Println("9开平方根", math.Sqrt(9))
	log.Println("10的3次方", math.Pow(10, 3))
	log.Println("27开3次方根", math.Pow(27, 0.3333333333333))
}

func BigRatAndBigInt() {
	amount,_ := new(big.Int).SetString("100000000000000000000000000001", 10)
	log.Println(amount)

	amountS := big.NewRat(3, 2) // 1.5
	amountB := new(big.Rat).SetFrac(amount, big.NewInt(1))

	m := new(big.Rat).Mul(amountB, amountS)
	r := m.Mul(m, big.NewRat(1, 1))
	log.Println(r.Num())
}

func UseE() {
	a := 1.0320388897123451e18
	log.Println(a)
}

func BigRatToFloat() {
	a := big.NewRat(1, 3)
	log.Println(a.Float64())
}