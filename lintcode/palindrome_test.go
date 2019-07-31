package lintcode

import (
	"fmt"
	"testing"
)

/*
描述
中文
English
给定一个字符串，判断其是否为一个回文串。只考虑字母和数字，忽略大小写。

你是否考虑过，字符串有可能是空字符串？这是面试过程中，面试官常常会问的问题。

在这个题目中，我们将空字符串判定为有效回文。

您在真实的面试中是否遇到过这个题？
样例
样例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
解释: "amanaplanacanalpanama"
样例 2:

输入: "race a car"
输出: false
解释: "raceacar"
挑战
O(n) 时间复杂度，且不占用额外空间。
*/

func TestPalindrome(t *testing.T) {
	s := "raceacar"
	t.Log(isPalindrome(s))
}

/**
 * @param s: A string
 * @return: Whether the string is a valid palindrome
 */
func isPalindrome(s string) bool {
	// write your code here
	s1 := cleanString(s)
	return judgePalindrome(s1)
}

func cleanString(src string) string {
	return src
}

func judgePalindrome(s string) bool {
	bytes := []byte(s)
	if s == "" || len(bytes) == 1 {
		return true
	}
	front := 0
	tail := len(bytes) - 1
	for front < tail {
		if bytes[front] != bytes[tail] {
			fmt.Println(bytes[front])
			return false
		}
		front++
		tail--
	}
	return true
}
