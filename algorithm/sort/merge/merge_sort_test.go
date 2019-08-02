package merge

import (
	"fmt"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	arr := []int{3, 1, 2, 5, 6, 43, 4}
	fmt.Println(MergeSort(arr))
}
