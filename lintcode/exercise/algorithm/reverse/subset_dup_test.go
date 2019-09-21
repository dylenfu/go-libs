package reverse

import (
	"testing"
)

/*
18. 子集 II
中文English
给定一个可能具有重复数字的列表，返回其所有可能的子集。

样例
样例 1：

输入：[0]
输出：
[
  [],
  [0]
]
样例 2：

输入：[1,2,2]
输出：
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
挑战
你可以同时用递归与非递归的方式解决么？

注意事项
子集中的每个元素都是非降序的
两个子集间的顺序是无关紧要的
解集中不能包含重复子集
*/

// if 语句的含义是，
// 在 A[index:] 中，每个数字只能附着到 temp 中一次
// 判断方法是 A[i] != A[i-1]
// 但是 A[index] == A[index-1] 也没有关系
// 因为 A[index-1] 不在 A[index:] 中
// 而且，需要执行 A[i]!=A[i-1] 时，
// 可以肯定 i>=1，所以，不需要验证 i-1>=0

/**
 * @param nums: A set of numbers.
 * @return: A list of lists. All valid subsets.
 */
func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	cur := []int{}
	for i := 0; i <= len(nums); i++ {
		dfs(ans, nums, i, 0, cur)
	}
	return ans
}

func dfs(ans [][]int, nums []int, n, s int, cur []int) {
	if len(cur) == n {
		tmp := make([]int, n)
		copy(tmp, cur)
		ans = append(ans, tmp)
		return
	}
	for i := s; i < len(nums); i++ {
		cur = append(cur, nums[i])
		dfs(ans, nums, n, i+1, cur)
		cur = cur[:len(cur)-1]
	}
}

func TestSubsetsWithDup(t *testing.T) {
	a1 := []int{0}
	src := subsetsWithDup(a1)
	if !containSubset(src, []int{}) || !containSubset(src, []int{0}) {
		t.Fatal("[]int{0} do not contain []int{} or []int{0}")
	}

	a2 := []int{1, 2, 2}
	src = subsetsWithDup(a2)
	if !containSubset(src, []int{2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{2}")
	}
	if !containSubset(src, []int{1}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1}")
	}
	if !containSubset(src, []int{1, 2, 2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1,2,2}")
	}
	if !containSubset(src, []int{2, 2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{2,2}")
	}
	if !containSubset(src, []int{1, 2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1,2}")
	}
	if !containSubset(src, []int{}) {
		t.Fatal("[]int{1,2,2} do not contain []int{}")
	}
}

// go test -v github.com/dylenfu/go-libs/lintcode/exercise/algorithm/backtracking -run TestContainSubset
func TestContainSubset(t *testing.T) {
	src := [][]int{[]int{}, []int{0, 1}, []int{0, 2}, []int{1, 2}}
	dst := []int{0, 3}
	if containSubset(src, dst) == true {
		t.Fatal("[][]int{[]int{}, []int{0,1}, []int{0, 2}, []int{1,2}} should not contain []int{0,3}")
	}

	dst = []int{0, 1}
	if containSubset(src, dst) == false {
		t.Fatal("[][]int{[]int{}, []int{0,1}, []int{0, 2}, []int{1,2}} should contain []int{0,1}")
	}
}

func containSubset(src [][]int, dst []int) bool {
	if len(dst) == 0 {
		return true
	}

	for _, s := range src {
		is := true
		for _, d := range dst {
			if !contains(s, d) {
				is = false
				break
			}
		}
		if is == true {
			return true
		}
	}

	return false
}

func contains(src []int, dst int) bool {
	for _, v := range src {
		if v == dst {
			return true
		}
	}
	return false
}

// go test -v github.com/dylenfu/go-libs/lintcode/exercise/algorithm/backtracking -run TestContains
func TestContains(t *testing.T) {
	src := []int{1, 2, 3}
	if contains(src, 3) == false {
		t.Fatal("[]int{1,2,3} should contain 3")
	}
	if contains(src, 0) == true {
		t.Fatal("[]int{1,2,3} should not contain 0")
	}
}
