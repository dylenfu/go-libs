package lintcode

import (
	"testing"
	"time"
)

/*
描述
中文
English
给定一个二叉树，判断它是否是合法的二叉查找树(BST)

一棵BST定义为：

节点的左子树中的值要严格小于该节点的值。
节点的右子树中的值要严格大于该节点的值。
左右子树也必须是二叉查找树。
一个节点的树也是二叉查找树。
您在真实的面试中是否遇到过这个题？
样例
样例 1:

输入：{-1}
输出：true
解释：
二叉树如下(仅有一个节点）:
	-1
这是二叉查找树。
样例 2:

输入：{2,1,4,#,#,3,5}
输出：true
解释：
	二叉树如下：
	  2
	 / \
	1   4
	   / \
	  3   5
这是二叉查找树。
*

/**
 * @param root: The root of binary tree.
 * @return: True if the binary tree is BST, or false
*/
func isValidBST(root *TreeNode) bool {
	// write your code here
	return false
}

func TestStackLimit(t *testing.T) {
	s := []string{}
	for i := 0; i < 100000000; i++ {
		s = append(s, "abcdef")
	}
	t.Log(len(s))
}

func TestStackPrint(t *testing.T) {
	t.Log(time.Now())
	t.Log(time.Now().Unix())
}
