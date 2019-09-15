package quicksort

import (
	"testing"
)

/*
给定一个未排序的整数数组，找到其中位数。

中位数是排序后数组的中间值，如果数组的个数是偶数个，则返回排序后数组的第N/2个数。

样例
样例 1:

输入：[4, 5, 1, 2, 3]
输出：3
解释：
经过排序，得到数组[1,2,3,4,5]，中间数字为3
样例 2:

输入：[7, 9, 4, 5]
输出：5
解释：
经过排序，得到数组[4,5,7,9]，第二个(4/2)数字为5
挑战
时间复杂度为O(n)

注意事项
数组大小不超过10000
*/

/*
解题思路:这里我们练习的是快速排序
快排思想:
快速排序采用分治法(devide and conquer), 选取一个参考值(privot)比如slice[0],
在[1, len-1]范围内扫描，head从左到右，tail从右到左，
小于privot的值放到数组左边(与head-坐标为0-进行交换)，
大于privot的值放到数组右边(与tail-坐标为len-1-进行交换)
最终将数组分成两部分,
然后将原数组slice[:head], slice[head+1:]部分分别进行递归。

需要注意几点:
1.数组元素小于等于1时无需再排序
2.privot,head的选为0, 而i的起始坐标是1, 这样可以减少一次比较
3.i的终止坐标是len-1.不要写成i < len - 1,这样会漏掉最后一个元素与privot的比较
4.i的走向与head一致，head对应的num小于privot时，head++，同时i++
5.拆分成两个数组时，不要把参考值坐标也放进去(nums[:head], nums[head:])
*/
func median(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return nums[0]
	}

	qsort(nums)
	mid := length / 2
	if length%2 == 0 {
		mid -= 1
	}
	return nums[mid]
}

// 默认长度大于1
func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	privot := nums[0] // idx =0 or random or other
	head := 0
	tail := len(nums) - 1

	for i := 1; i <= tail; {
		if nums[i] > privot {
			swap(nums, tail, i)
			tail--
		} else {
			swap(nums, head, i)
			head++
			i++
		}
	}

	qsort(nums[:head])
	qsort(nums[head+1:])
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// go test -v github.com/dylenfu/go-libs/lintcode/exercise/algorithm/quicksort -run TestMedian
func TestMedian(t *testing.T) {
	if x := median([]int{4, 5, 1, 2, 3}); x != 3 {
		t.Log(x)
		t.Log("median([]int{4, 5, 1, 2, 3}) != 3")
	}
	if x := median([]int{7, 9, 4, 5}); x != 5 {
		t.Log(x)
		t.Fatal("median([]int{7, 9, 4, 5}) != 5")
	}
	if x := median([]int{1, 0, 1}); x != 1 {
		t.Log(x)
		t.Fatal("median([]int{1, 0, 1} != 1")
	}
}
