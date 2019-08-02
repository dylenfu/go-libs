package skiplist

import (
	"math/rand"
	"testing"
	"time"
)

func TestSkipList(t *testing.T) {
	list := New()
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 1000; i++ {
		//if i == 6 {
		//	fmt.Println("debug here")
		//}
		x := rand.Int31()
		list.Insert(int(x), nil)
	}
	list.PrintList()
}
