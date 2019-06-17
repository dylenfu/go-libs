package lintcode

import (
	"fmt"
	"testing"
)

func TestTriangleSelect(t *testing.T) {
	arr := []int{2,3,5,8}
	t.Log(judgeTriangle(arr))
}

func judgeTriangle(arr []int) string {
	ret := "no"
	if len(arr) < 3 {
		return ret
	}
	cnt := 0
	length := len(arr)
	for i:=0; i<length-2;i++ {
		for j:=i+1; j<length-1;j++ {
			for k:=j+1; k<length;k++ {
				cnt++
				fmt.Println(arr[i], arr[j], arr[k])
				if judgeUnit(arr[i], arr[j], arr[k]) {
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

func judgeUnit(a, b, c int) bool {
	ret := false
	if a + b > c && a +  c > b && b + c > a {
		ret = true
	}
	return ret
}