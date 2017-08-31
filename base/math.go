package base

import (
	"math"
	"log"
)

func SimpleMath() {
	log.Println("9开平方根",math.Sqrt(9))
	log.Println("10的3次方", math.Pow(10,3))
	log.Println("27开3次方根", math.Pow(27,0.3333333333333))
}