package base

import (
	"testing"
	"time"
)

func TestSimpleGoRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			println(i)
		}(i)
	}
	time.Sleep(10 * time.Second)
}
