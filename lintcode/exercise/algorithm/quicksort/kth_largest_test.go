package quicksort

import "testing"

/*
5. 第k大元素
中文English
在数组中找到第 k 大的元素。

样例
样例 1：

输入：
n = 1, nums = [1,3,4,2]
输出：
4
样例 2：

输入：
n = 3, nums = [9,3,2,4,8]
输出：
4
挑战
要求时间复杂度为O(n)，空间复杂度为O(1)。

注意事项
你可以交换数组中的元素的位置
*/

/*
解题思路:
快排时需要将数组拆分成两个数组，左边的大于privot，右边的小于privot(desc)。
如果刚好要寻找的第n个元素就是分界线(后面称为head)，那么我们就可以直接返回了。
如果n跟head不一致，那么需要判断要查询的这个元素在那一边。
如果是在左边,n=3,要查的是第3大的数据，那么就在原数组内继续查询，
因为到目前为止，还没有确定左边区间内有任何一个元素是可以放弃的，如下:
9    8    7    6
------------------------------
0        n-1   head         ...

如果是在右边, n=5, 要查的是第5大的数据， 那么我们应该在head后(不包含head)查询第n-(head+1)大的元素
这里，没有必要纠结n-head-1到底怎么来的，只需要简单地想，包含head在内，前面head+1个元素已经直接被舍弃了
9    8    7    6    5
------------------------------
0        head      n -1     ...

如果刚好n - 1 == head就直接返回

*/

func kthLargestElement(n int, nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	head := 0
	tail := len(nums) - 1
	privot := nums[0]

	for i := 1; i <= tail; {
		if nums[i] > privot {
			swap(nums, i, head)
			head++
			i++
		} else {
			swap(nums, i, tail)
			tail--
		}
	}

	if head == n-1 {
		return nums[head]
	} else if head < n-1 {
		return kthLargestElement(n-head-1, nums[head+1:])
	} else {
		return kthLargestElement(n, nums[:head])
	}
}

// go test -v github.com/dylenfu/go-libs/lintcode/exercise/algorithm/quicksort -run TestKthLargestElement
func TestKthLargestElement(t *testing.T) {
	if x := kthLargestElement(1, []int{1, 3, 4, 2}); x != 4 {
		t.Log(x)
		t.Fatal("[]int{1,3,4,2} 1st element is not 4")
	}

	if x := kthLargestElement(3, []int{9, 3, 2, 4, 8}); x != 4 {
		t.Log(x)
		t.Fatal("[]int{9,3,2,4,8} 3rd element is not 4")
	}
}
