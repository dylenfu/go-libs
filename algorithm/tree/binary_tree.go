package tree

import (
	"fmt"
)

type UserExpEntry struct {
	left, right *UserExpEntry
	value       uint64 // exp 低32位 & value 高32位
}

func NewEntry(uid int32, exp uint32, left, right *UserExpEntry) (entry *UserExpEntry) {
	entry = &UserExpEntry{value: formatValue(uid, exp), left: left, right: right}
	return
}

func (root *UserExpEntry) Insert(a *UserExpEntry) {
	for {
		if a.value > root.value {
			if root.right == nil {
				root.right = a
				return
			}
			root = root.right

		} else if a.value < root.value {
			if root.left == nil {
				root.left = a
				return
			}
			root = root.left
		} else {
			return
		}
	}
}

func (root *UserExpEntry) Find(exp uint64) (result *UserExpEntry) {
	for {
		if exp > root.value {
			root = root.right
			if root == nil {
				break
			}
		} else if exp < root.value {
			root = root.left
			if root == nil {
				break
			}
		} else {
			result = root
			break
		}
	}
	return
}

func (a *UserExpEntry) FindPrint(exp uint64) {
	result := a.Find(exp)
	if result != nil {
		fmt.Println(result.value)
	}
}

//先序遍历
func (a *UserExpEntry) LeftPrint() {
	if a.left != nil {
		l := a.left
		l.LeftPrint()
	}
	fmt.Print("  ")
	fmt.Print(a.value)
	if a.right != nil {
		r := a.right
		r.LeftPrint()
	}
}

func (root *UserExpEntry) FindBefore(exp uint64) *UserExpEntry {
	var result *UserExpEntry
	if exp == root.value {
		fmt.Println("不能移除根节点")
		return nil
	}
	for true {
		if (root.right != nil && exp == root.right.value) || (root.left != nil && exp == root.left.value) {
			return root
		} else if exp > root.value {
			root = root.right
			if root == nil {
				fmt.Println("不存在")
				break
			}
		} else if exp < root.value {
			root = root.left
			if root == nil {
				fmt.Println("不存在")
				break
			}
		} else if root.left == nil && root.right == nil {
			fmt.Println("不存在")
			break
		}
	}
	return result
}

//删除元素
func (a *UserExpEntry) Remove(v uint64) {
	rm := a.Find(v)
	be := a.FindBefore(v)
	if rm != nil && be != nil {
		//要删除的元素的左右节点都为空
		if rm.left == nil && rm.right == nil {
			if rm.value > be.value {
				be.right = nil
			} else {
				be.left = nil
			}
		} else if rm.left == nil && rm.right != nil {
			//左节点为空 右节点不为空
			if rm.value > be.value {
				be.right = rm.right
			} else {
				be.left = rm.right
			}

		} else if rm.left != nil && rm.right == nil {
			//左节点不为空 右节点为空
			if rm.value > be.value {
				be.right = rm.left
			} else {
				be.left = rm.left
			}

		} else {
			//左节点不为空 右节点不为空

			//把前驱节点的指针断开
			if rm.value > be.value {
				be.right = nil
			} else {
				be.left = nil
			}
			//把rm.left及其下面的节点 以 rm.left 为代表作为一个整体重新插入
			a.Insert(rm.left)
			a.Insert(rm.right)
		}
	}

}
