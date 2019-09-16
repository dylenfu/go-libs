package quicksort

import "testing"

/*
508. 摆动排序
中文English
给你一个没有排序的数组，请将原数组就地重新排列满足如下性质

nums[0] <= nums[1] >= nums[2] <= nums[3]....
样例
样例 1:

输入: [3, 5, 2, 1, 6, 4]
输出: [1, 6, 2, 5, 3, 4]
解释: 这个问题可能有多种答案, [2, 6, 1, 5, 3, 4] 同样可以.
样例 2:

输入: [1, 2, 3, 4]
输出: [1, 4, 2, 3]
注意事项
请就地排序数组，不要额外定义数组.
*/

/*
思路
根据快排思想，取到某个中间值，比他小的一半和比他大的一半进行穿插，则可以实现这种摆动
比如说 1, 2, 3, 4, 5, 6, 7 中间值是4, 比中间值小的一半是(1,2,3),比中间值大的一半是(5,6,7)
穿插后是(1, 5, 2, 6, 3, 7, 5)
另一种情况，(1,2,3,4,5,6) 中间值是3, 比他小的一半是(1,2), 比他大的一半是(4,5,6)
穿插时可以用数量多的一半放前面(4, 1, 5, 2, 6, 3)

那么，步骤就是先找到这个中间值，然后进行穿插
*/
func wiggleSort(nums []int) []int {
	midIdx := getMidIdx(nums)
	qsortForMid(nums, midIdx)
	return mergeForMid(nums, midIdx)
}

func qsortForMid(nums []int, midIdx int) {
	if len(nums) <= 1 {
		return
	}

	privot := nums[0]
	head := 0
	tail := len(nums) - 1

	for i := 1; i <= tail; {
		if nums[i] > privot {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}

	if head == midIdx {
		return
	} else if head > midIdx {
		qsortForMid(nums[:head], midIdx)
	} else {
		qsortForMid(nums[head+1:], midIdx-head-1)
	}
}

func getMidIdx(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return 0
	}

	if length%2 == 0 {
		return length/2 - 1
	} else {
		return length / 2
	}
}

func mergeForMid(nums []int, mid int) []int {
	var (
		nums1, nums2 []int
		i1, i2, i    = 0, 0, 0
	)

	length := len(nums)
	ret := make([]int, length)

	if length/2 == mid {
		nums1 = nums[:mid]
		nums2 = nums[mid+1:]
	} else {
		nums1 = nums[mid+1:]
		nums2 = nums[:mid]
	}

	for i1 < len(nums1) && i2 < len(nums2) {
		ret[i], ret[i+1] = nums1[i1], nums2[i2]
		i += 2
		i1++
		i2++
	}
	if i1 < len(nums1) {
		ret[i] = nums1[i1]
	} else if i2 < len(nums2) {
		ret[i] = nums2[i2]
	}
	ret[length-1] = nums[mid]

	return ret
}

// go test -v github.com/dylenfu/go-libs/lintcode/exercise/algorithm/quicksort -run TestWiggleSort
func TestWiggleSort(t *testing.T) {
	nums1 := []int{3, 5, 2, 1, 6, 4}
	ret1 := wiggleSort(nums1)
	if !validateWiggleSort(ret1) {
		t.Fatal("[]int{3, 5, 2, 1, 6, 4} after sorted is", nums1)
	}

	nums2 := []int{1, 2, 3, 4}
	ret2 := wiggleSort(nums2)
	if !validateWiggleSort(ret2) {
		t.Fatal("[]int{1, 2, 3, 4} after sorted is", nums2)
	}
}

func validateWiggleSort(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	bigger := false
	if nums[1] > nums[0] {
		bigger = true
	}

	for i := 2; i < len(nums); i++ {
		if bigger == true {
			if nums[i] > nums[i-1] {
				return false
			} else {
				bigger = false
			}
		} else {
			if nums[i] < nums[i-1] {
				return false
			} else {
				bigger = true
			}
		}
	}

	return true
}
