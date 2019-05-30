package bubble

import (
	"testing"
	"fmt"
)

func Test_BubbleSort(t *testing.T) {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27,12}
	BubbleSort(values)
	fmt.Println(values)
}
