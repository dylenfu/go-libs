package base

import (
	"fmt"
	"testing"
)

func TestCompactError(t *testing.T) {
	var err error

	wr := func(i int) {
		if err != nil {
			return
		}
		err = write(i)
	}

	wr(1)
	wr(2)
	wr(3)
	wr(4)

	if err != nil {
		t.Log(err)
	}
}

// 奇数时error为nil，偶数时不为空
func write(i int) error {
	fmt.Println(i)
	if i % 2 == 0 {
		return nil
	}
	return fmt.Errorf("error %d", i)
}