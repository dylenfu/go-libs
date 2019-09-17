package mergesort

import (
	"testing"
)

/*
532. 逆序对
中文English
在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。给你一个数组，求出这个数组中逆序对的总数。
概括：如果a[i] > a[j] 且 i < j， a[i] 和 a[j] 构成一个逆序对。

样例
样例1

输入: A = [2, 4, 1, 3, 5]
输出: 3
解释:
(2, 1), (4, 1), (4, 3) 是逆序对
样例2

输入: A = [1, 2, 3, 4]
输出: 0
解释:
没有逆序对
*/

/*
思路:
归并排序通过分治的方式实现排序，
{3, 1, 21, 4, 5}, 拆分成{3, 1}, {21, 4, 5}
{3, 1} 进一步拆分成{3}, {1}, 合并时3 > 1会产生一个逆序
{21, 4, 5} 进一步拆分成{21}, {4, 5}, 其中{4, 5} 进一步拆分, 排序时4 < 5 不形成逆序
而{21} , {4, 5}合并时会有两次判断并生成逆序

简单的说，就是归并排序过程中合并时如果i < j && a[i] > a[j]的判断可以生成一个逆序对
需要注意的是,merge时还是需要将数据append到数组，因为递归的原因，新生成的数组需要参与后续的归并
*/
func reversePairs(A []int) int64 {
	var cnt int64 = 0
	mergeSort(A, &cnt)
	return cnt
}

func mergeSort(nums []int, n *int64) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSort(nums[:mid], n)
	right := mergeSort(nums[mid:], n)

	return merge(left, right, n)
}

// merge
func merge(left, right []int, n *int64) (result []int) {
	l, r := 0, 0
	leftswitch := true
	leftinc := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
			leftswitch = true
		} else {
			result = append(result, right[r])
			r++
			if leftswitch {
				*n += int64(r)
				leftswitch = false
				leftinc++
			} else {
				*n += 1
			}
		}
	}

	if remain := len(left[l:]); remain != 0 {
		result = append(result, left[l:]...)
		if remain > leftinc {
			*n += int64((remain - leftinc) * r)
		}
	} else {
		result = append(result, right[r:]...)
	}

	return
}

// go test -v github.com/dylenfu/go-libs/lintcode/exercise/algorithm/mergesort -run TestReversePairs
func TestReversePairs(t *testing.T) {
	if x := reversePairs([]int{2, 4, 1, 3, 5}); x != 3 {
		t.Log(x)
		t.Fatal("reversePairs([]int{2, 4, 1, 3, 5}) != 3")
	}

	if x := reversePairs([]int{1, 2, 3, 4}); x != 0 {
		t.Log(x)
		t.Fatal("reversePairs([]int{1, 2, 3, 4}) != 0")
	}

	if x := reversePairs([]int{4, 3, 2, 1}); x != 6 {
		t.Log(x)
		t.Fatal("reversePairs([]int{4,3,2,1}) != 6")
	}

	if x := reversePairs([]int{2, 3, 1, 55, 6, 4, 7, 3, 0}); x != 18 {
		t.Log(x)
		t.Fatal("reversePairs([]int{2,3,1,55,6,4,7,3,0}) != 18")
	}
}
