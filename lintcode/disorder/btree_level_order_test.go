package disorder

import "testing"

/*
样例
样例 1:

输入：{1,2,3}
输出：[[1],[2,3]]
解释：
   1
  / \
 2   3
它将被序列化为{1,2,3}

层次遍历
样例 2:

输入：{1,#,2,3}
输出：[[1],[2],[3]]
解释：
1
 \
  2
 /
3
它将被序列化为{1,#,2,3}
层次遍历
挑战
挑战1：只使用一个队列去实现它

挑战2：用BFS算法来做

注意事项
首个数据为根节点，后面接着是其左儿子和右儿子节点值，"#"表示不存在该子节点。
节点数量不超过20。
*/
var ret [][]int

/**
 * @param root: A Tree
 * @return: Level order a list of lists of integer
 */
func levelOrder(root *TreeNode) [][]int {
	// write your code here
	levelOrderHelper(root, 0)
	return ret
}

func levelOrderHelper(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(ret) < level+1 {
		ret = append(ret, make([]int, 0))
	}

	data := ret[level]
	data = append(data, root.Val)
	ret[level] = data

	levelOrderHelper(root.Left, level+1)
	levelOrderHelper(root.Right, level+1)
}

func TestLevelOrder1(t *testing.T) {
	node3 := TreeNode{Val: 3, Left: nil, Right: nil}
	node2 := TreeNode{Val: 2, Left: nil, Right: nil}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}

	orders := levelOrder(&node1)
	t.Log(orders)
}

func TestLevelOrder2(t *testing.T) {
	node6 := TreeNode{Val: 6, Left: nil, Right: nil}
	node5 := TreeNode{Val: 5, Left: nil, Right: nil}
	node4 := TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := TreeNode{Val: 3, Left: &node6, Right: nil}
	node2 := TreeNode{Val: 2, Left: &node4, Right: &node5}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}

	orders := levelOrder(&node1)
	t.Log(orders)
}
