package lintcode

/*
描述
中文
English
给定一个数组arr，问，能否从数组里找到3个元素作为三条边的边长，使三条边能够组成一个三角形。若能，返回yes,若不能，返回no

1 \leq n \leq 1000001≤n≤100000
1 \leq arr[i] \leq 10000000001≤arr[i]≤1000000000
程序会被运行500次
您在真实的面试中是否遇到过这个题？
样例
样例 1:

输入：arr=[2,3,5,8]
输出："no"
解释：
2,3,5无法组成三角形
2,3,8无法组成三角形
3,5,8无法组成三角形
所以，返回"no"
样例 2:

输入：arr=[3,4,5,8]
输出："yes"
解释：
3,4,5可以组成一个三角形
所以返回"yes"
*/

import (
	"fmt"
	"testing"
)

func TestTriangleSelect(t *testing.T) {
	arr := []int{2, 3, 5, 8}
	t.Log(judgeTriangle(arr))
}

func judgeTriangle(arr []int) string {
	ret := "no"
	if len(arr) < 3 {
		return ret
	}
	cnt := 0
	length := len(arr)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				cnt++
				fmt.Println(arr[i], arr[j], arr[k])
				if judgeTriangleUnit(arr[i], arr[j], arr[k]) {
					ret = "yes"
					return ret
				}
				if cnt > 500 {
					return ret
				}
			}
		}
	}

	return ret
}

func judgeTriangleUnit(a, b, c int) bool {
	ret := false
	if a+b > c && a+c > b && b+c > a {
		ret = true
	}
	return ret
}
