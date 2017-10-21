package base

import "time"

func BreakLoop() {
	for i:=0;i<10;i++ {
	next1:
		for j:=0;j<10;j++ {
			if j == 2 {
				println(i, j)
				break next1
			}
		}
	}
}

func BreakForLoop() {
	i := 0
	for {
		if simple(i) {
			i++
			println("========11111")
			time.Sleep(1 * time.Second)
			continue
		}

		for j:=0;j<10;j++ {
			if j < 5 {
				continue
			} else {
				time.Sleep(1 * time.Second)
				for k := 0; k< 10;k++ {
					if k < 5 {
						continue
					}

					println("========22222")
				}
			}
		}

		println("=========33333")

	}
}

func simple(i int) bool {
	if i > 10 {
		return false
	} else {
		return true
	}
}