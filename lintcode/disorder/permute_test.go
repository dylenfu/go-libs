package disorder

import "testing"

/*
描述
中文
English
给定一个数字列表，返回其所有可能的排列。

你可以假设没有重复数字。

您在真实的面试中是否遇到过这个题？
样例
样例 1：

输入：[1]
输出：
[
[1]
]
样例 2：

输入：[1,2,3]
输出：
[
[1,2,3],
[1,3,2],
[2,1,3],
[2,3,1],
[3,1,2],
[3,2,1]
]
挑战
使用递归和非递归分别解决。
*/

func TestPermute(t *testing.T) {
	arr := []int{}
	permute(arr)
}

func permute(arr []int) {

}
