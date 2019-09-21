package reverse

import "testing"

/*
18. 子集 II
中文English
给定一个没有重复数字的列表，返回其所有可能的子集。

样例
样例 1：

输入：[0]
输出：
[
  [],
  [0]
]
样例 2：

输入：[1,2,3]
输出：
[
  [1],
  [2],
  [3],
  [1,2]
  [1,3]
  [2,3]
  [1,2,3],
  []
]
*/

/*
解题思路:
包含n个元素的数组含有2^n个子集(含重复内容)，可以根据元素在数组中所在位置表示为2进制数，
比如[1,2,3]这样的数组含1,2,2总共3个元素
000, 001, 010, 011, 100, 101, 110, 111对应为
{}   {3}  {2} {2,3} {1} {1,3} {1,2} {1,2,3}
*/

/**
 * @param nums: A set of numbers.
 * @return: A list of lists. All valid subsets.
 */
func subsetsWithDup1(nums []int) [][]int {
	empty := []int{}
	n := len(nums)
	if n == 0 {
		return [][]int{empty}
	}

	ret := [][]int{empty}

	// i代表有i个元素
	max := 1 << uint(n)
	eles := make([]int, n, n)
	for i := 1; i <= max-1; i++ {
		binarayInc(eles, 1, 0)
		data := []int{}
		for j := 0; j < n; j++ {
			if eles[j] == 1 {
				data = append(data, nums[j])
			}
		}
		ret = append(ret, data)
	}
	return ret
}

/*
解题思路:
遍历nums数组，同时遍历res数组，将res之前的元素内容添加上新的元素
e.g
nums = {1,2,3}
n = 1 => 新增 append({1}, nil)  => res{ {1} }
n = 2 => 新增 append({2}, nil), append({2}, {1}...) => res { {1}, {2}, {1,2} }
n = 3 => 新增 append({3}, nil), append({3}, {1}...), append({3}, {2}...), append({3}, {1,2}...)
所以最后有{1}, {2}, {1,2}, {3}, {3,1}, {3,2}, {3,1,2}
*/
func subsetsWithDup2(nums []int) [][]int {
	res := make([][]int, 1, 1024)
	for _, n := range nums {
		for _, r := range res {
			res = append(res, append([]int{n}, r...))
		}
	}
	return res
}

func binarayInc(slice []int, inc int, index int) {
	m := slice[index] + inc
	if m < 2 {
		slice[index] = m
		return
	} else {
		slice[index] = 0
		binarayInc(slice, 1, index+1)
	}
}

func TestSubsetsWithDup1(t *testing.T) {
	a1 := []int{0}
	src := subsetsWithDup1(a1)
	if !containSubset(src, []int{}) || !containSubset(src, []int{0}) {
		t.Fatal("[]int{0} do not contain []int{} or []int{0}")
	}

	a2 := []int{1, 2, 3}
	src = subsetsWithDup1(a2)
	if !containSubset(src, []int{1}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1}")
	}
	if !containSubset(src, []int{2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{2}")
	}
	if !containSubset(src, []int{3}) {
		t.Fatal("[]int{1,2,2} do not contain []int{3}")
	}
	if !containSubset(src, []int{1, 2}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1,2}")
	}
	if !containSubset(src, []int{1, 3}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1,3}")
	}
	if !containSubset(src, []int{2, 3}) {
		t.Fatal("[]int{1,2,2} do not contain []int{2,3}")
	}
	if !containSubset(src, []int{1, 2, 3}) {
		t.Fatal("[]int{1,2,2} do not contain []int{1,2,3}")
	}
	if !containSubset(src, []int{}) {
		t.Fatal("[]int{1,2,2} do not contain []int{}")
	}
}
