package quicksort

import "testing"

/*
整数排序 II
中文English
给一组整数，请将其在原地按照升序排序。使用归并排序，快速排序，堆排序或者任何其他 O(n log n) 的排序算法。

样例
例1：

输入：[3,2,1,4,5]，
输出：[1,2,3,4,5]。
例2：

输入：[2,3,1]，
输出：[1,2,3]。
*/

/*
解题思路，参考median_test
*/
func sortIntegers2(A *[]int) {
	if len(*A) <= 1 {
		return
	}

	head := 0
	tail := len(*A) - 1
	privot := (*A)[0]

	for i := 1; i <= tail; {
		if (*A)[i] > privot {
			swapA(A, i, tail)
			tail--
		} else {
			swapA(A, i, head)
			head++
			i++
		}
	}

	x := (*A)[:head]
	y := (*A)[head+1:]
	sortIntegers2(&x)
	sortIntegers2(&y)
}

func swapA(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

// go test -v github.com/dylenfu/go-libs/lintcode/exercise/algorithm/quicksort -run TestQSort2
func TestQSort2(t *testing.T) {
	nums1 := []int{3, 2, 1, 4, 5}
	sortIntegers2(&nums1)
	if nums1[0] != 1 || nums1[1] != 2 || nums1[2] != 3 || nums1[3] != 4 || nums1[4] != 5 {
		t.Fatal("sortIntegers2([]int{3,2,1,4,5}) != []int{1,2,3,4,5}")
	}

	nums2 := []int{2, 3, 1}
	sortIntegers2(&nums2)
	if nums2[0] != 1 || nums2[1] != 2 || nums2[2] != 3 {
		t.Fatal("sortIntegers2([]int{2,3,1}) != []int{1,2,3}")
	}

	nums3 := []int{}
	sortIntegers2(&nums3)
	t.Log(nums3)
	if len(nums3) != 0 {
		t.Fatal("len(nums3) != 0")
	}
}
