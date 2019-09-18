package backtracking

import "testing"

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

/**
 * @param nums: A set of numbers.
 * @return: A list of lists. All valid subsets.
 */
func subsetsWithDup (nums []int) [][]int {
	// write your code here
	return nil
}

func TestSubsetsWithDup(t *testing.T) {
	a1 := []int{0}
	src := subsetsWithDup(a1)
	if !containSubset(src, []int{}) || !containSubset(src, []int{1}) {
		t.Fatal("[]int{0} do not contain []int{} or []int{1}")
	}

	a2 := []int{1,2,2}
	src = subsetsWithDup(a2)
	if !containSubset(src, []int{2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{2}")
	}
	if !containSubset(src, []int{1}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1}")
	}
	if !containSubset(src, []int{1,2,2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1,2,2}")
	}
	if !containSubset(src, []int{2,2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{2,2}")
	}
	if !containSubset(src, []int{1,2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1,2}")
	}
	if !containSubset(src, []int{}) {
		t.Fatal("[]int{1,2,2} do not contain []int{}")
	}
}

// go test -v github.com/dylenfu/go-libs/lintcode/exercise/algorithm/backtracking -run TestContainSubset
func TestContainSubset(t *testing.T) {
	src := [][]int{[]int{}, []int{0,1}, []int{0, 2}, []int{1,2}}
	dst := []int{0,3}
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
	src := []int{1,2,3}
	if contains(src, 3) == false {
		t.Fatal("[]int{1,2,3} should contain 3")
	}
	if contains(src, 0) == true {
		t.Fatal("[]int{1,2,3} should not contain 0")
	}
}