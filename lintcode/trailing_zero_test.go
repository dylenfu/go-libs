package lintcode

import (
	"testing"
)

/*
2. 尾部的零
中文English
设计一个算法，计算出n阶乘中尾部零的个数

Example
样例  1:
	输入: 11
	输出: 2

	样例解释:
	11! = 39916800, 结尾的0有2个。

样例 2:
	输入:  5
	输出: 1

	样例解释:
	5! = 120， 结尾的0有1个。

Challenge
O(logN)的时间复杂度
*/

/*
该题来自《编程之美》，n的阶乘可以理解成质数相乘的形式:
1 * 2 * 3 *     4   * 5 *    6    * 7 *     8      *    9    *   10 .. * n -->
1 * 2 * 3 * (2 * 2) * 5 * (2 * 3) * 7 *(2 * 2 * 2) * (3 * 3) *(2 * 5) *....
尾部有多少0，就意味着有多少次质数相乘等于10，质数相乘等于10只有1种情况 2 * 5
那么，质数2，5的个数的最小值min(2, 5)就是最终答案，2的跨度比5小很多，这样一来就是求5的个数，
包含质数5的数为5的倍数，或者5的x次方的倍数，
f(n) = n/5 + n/(5^2) + n/(5^3) +.....
*/

/**
 * @param n: A long integer
 * @return: An integer, denote the number of trailing zeros in n!
 * @desc: num为5的i次方
 */
func trailingZeros(n int64) int64 {
	var m, cnt int64 = 1, 0
	for {
		if m = m * 5; m > n {
			break
		}
		cnt += n / m
	}
	return cnt
}

// go test -v github.com/dylenfu/go-libs/lintcode -run TestTrailingZero
func TestTrailingZero(t *testing.T) {
	x := trailingZeros(105)
	if x != 25 {
		t.Log(x)
		t.Fatal("trailingZeros(10) ! = 1")
	}
}
