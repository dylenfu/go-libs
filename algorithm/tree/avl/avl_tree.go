package avl

import "sync"

/*
AVL树是一种平衡搜索树，具有以下特点
1.avl是二叉树，每个节点最多有两个子节点
2.左子树所有节点小于根节点
3.右子树所有节点大于根节点
4.某个节点的左右子树高度差不大于1

AVL树在查询效率上，时间复杂度为O(log2N))，与折半查找相当，
一般来说二叉树排序容易造成的左右子树有交大高度差，甚至是
线性查询，其时间复杂度接近为o(n), avl树能平衡高度差
*/

type AVLTree struct {
	head *AVLTreeNode
	mtx  *sync.RWMutex
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		mtx: new(sync.RWMutex),
	}
}

func (t *AVLTree) Insert(value Comparable) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	x := t.head.insert(value)
	t.head = x
}

func (t *AVLTree) Delete(value Comparable) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	t.head = t.head.delete(value)
}

func (t *AVLTree) Find(value Comparable) *AVLTreeNode {
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	node := t.head

	for {
		switch value.Compare(value, node.value) {
		case 1:
			node = node.right
		case -1:
			node = node.left
		case 0:
			return node
		}
	}
}

// after traversal tree, desc
func (t *AVLTree) Scan() (list []*AVLTreeNode) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	return t.head.afterTraversal()
}

func (t *AVLTree) Range(min, max Comparable) (list []*AVLTreeNode) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	return t.head.scope(min, max)
}

type AVLTreeNode struct {
	value       Comparable
	height      int
	left, right *AVLTreeNode
}

// getHeight 返回高度，节点为nil时返回0
func (a *AVLTreeNode) getHeight() int {
	if a == nil {
		return 0
	}
	return a.height
}

// min 以a为根节点，寻找并返回该子树的最小节点
func (a *AVLTreeNode) min() *AVLTreeNode {
	if a == nil {
		return nil
	}
	if a.left == nil {
		return a
	}
	return a.left.min()
}

// max 以a为根节点，寻找该子树的最大节点 返回新的根节点
func (a *AVLTreeNode) max() *AVLTreeNode {
	if a == nil {
		return nil
	}
	if a.right == nil {
		return a
	}
	return a.right.max()
}

// updateHeight 更新高度，取左右子树高度最大值，然后递增1
func (a *AVLTreeNode) updateHeight() {
	a.height = maxint(a.left.getHeight(), a.right.getHeight()) + 1
}

// balanceFactor 返回左右子树的高度差
func (a *AVLTreeNode) balanceFactor() int {
	if a == nil {
		return 0
	}
	return a.left.getHeight() - a.right.getHeight()
}

/*
          x             y
           \           / \
           y     ->   x  t
            \
            t

         pic-1         pic-2

如图中pic-1所示，如果再继续在右子树插入右孩子，会导致AVL失衡
为了处理这种情况，需要将x节点左旋，降低右子树高度。然后再插入右孩子

leftRotate 左旋，返回该子树新的根节点(即pic-2中y)
*/
func (x *AVLTreeNode) leftRotate() *AVLTreeNode {
	y := x.right
	x.right = y.left
	y.left = x
	x.updateHeight()
	y.updateHeight()
	return y
}

/*
            x             y
           /             / \
          y        ->   t  x
         /
        t
          pic-1         pic-2

如图中pic-1所示，如果要继续往左子树插入左孩子，会导致AVL失衡
为了处理这种情况，需要将y右旋，降低左子树高度，然后再插入左孩子

rightRotate 右旋，返回该子树新的根节点(即pic-2中x)
*/
func (x *AVLTreeNode) rightRotate() *AVLTreeNode {
	y := x.left
	x.left = y.right
	y.right = x
	x.updateHeight()
	y.updateHeight()
	return y
}

/*
		x                x              z
       /                /              / \
      y      ->        z      ->      y   x
       \              /
        z            y

      pic-1           pic-2            pic-3

如图中pic-1所示 在x的左子树插入右孩子z，会导致失衡
1、先将x的右孩子进行左旋，y下去，z上来形成pic-2，这时，将操作得到的新根节点z重新给到x的左节点
2、然后再对x根节点进行右旋，z上去，x下来，形成pic-3

leftRightRotate 先左后右，具体解决往左子树插右孩子的具体情况
*/
func (a *AVLTreeNode) leftRightRotate() *AVLTreeNode {
	a.left = a.left.leftRotate()
	return a.rightRotate()
}

/*
         x	            x             z
          \              \           / \
          y    ->        z     ->   x   y
         /                \
        z                 y
       pic-1           pic-2         pic-3

如图中pic-1所示，往右子树插入左孩子会造成AVL失衡
1、先将x的右孩子进行右旋，y下去，z上来，形成pic-2，这时，将操作所得的根节点z给到x的右节点
2、然后在对x进行左旋，x下去，z上来，形成pic-3

rightLeftRotate 先右后左，解决往右子树插入左节点的场景
*/
func (a *AVLTreeNode) rightLeftRotate() *AVLTreeNode {
	a.right = a.right.rightRotate()
	return a.leftRotate()
}

// rebuildBalance 根据插入或者删除的值，以a为节点重建avl平衡 返回新的根节点
func (a *AVLTreeNode) rebuildBalance() (node *AVLTreeNode) {
	a.updateHeight()
	switch a.balanceFactor() {
	case 2:
		if a.left.balanceFactor() == -1 {
			node = a.leftRightRotate()
		} else {
			node = a.rightRotate()
		}
		break
	case -2:
		if a.right.balanceFactor() == 1 {
			node = a.rightLeftRotate()
		} else {
			node = a.leftRotate()
		}
		break
	default:
		node = a
	}
	return
}

// insert 以a为入口节点，插入新增节点 并返回新的根节点
func (a *AVLTreeNode) insert(value Comparable) *AVLTreeNode {
	if a == nil {
		return &AVLTreeNode{value: value, height: 1}
	}

	switch value.Compare(value, a.value) {
	case -1:
		a.left = a.left.insert(value)
		break
	case 0, 1:
		a.right = a.right.insert(value)
		break
	}
	return a.rebuildBalance()
}

// delete 以a为入口节点，删除某个数据对应的节点 并返回新的根节点
func (a *AVLTreeNode) delete(value Comparable) *AVLTreeNode {
	if a == nil {
		return nil
	}

	switch value.Compare(value, a.value) {
	case -1:
		a.left = a.left.delete(value)
	case 1:
		a.right = a.right.delete(value)
	case 0:
		// 这个时候还处于平衡状态，也就是说，左右子树高度差还小于2,
		// 当有左/右子树为空时可以直接返回另一半，不需要rebuildBalance
		// 如果左子树较高则取左子树最大值替代根节点, 否则取右子树最小值替代根节点
		if a.left != nil && a.right != nil {
			if a.left.getHeight() > a.right.getHeight() {
				max := a.left.max()
				a.value = max.value
				a.left = a.left.delete(max.value)
			} else {
				min := a.right.min()
				a.value = min.value
				a.right = a.right.delete(min.value)
			}
		} else if a.left != nil && a.right == nil {
			a = a.left
		} else if a.right != nil && a.left == nil {
			a = a.right
		} else {
			a = nil
		}

		if a != nil {
			a.rebuildBalance()
		}
	}

	return a
}

// after traversal tree, desc
func (a *AVLTreeNode) afterTraversal() (list []*AVLTreeNode) {
	if a == nil {
		return
	}

	if a.right != nil {
		list = append(list, a.right.afterTraversal()...)
	}
	list = append(list, a)
	if a.left != nil {
		list = append(list, a.left.afterTraversal()...)
	}
	return
}

func (a *AVLTreeNode) preTraversal() (list []*AVLTreeNode) {
	if a == nil {
		return
	}

	if a.left != nil {
		list = append(list, a.left.preTraversal()...)
	}
	list = append(list, a)
	if a.right != nil {
		list = append(list, a.right.preTraversal()...)
	}
	return
}

func (a *AVLTreeNode) scope(min, max Comparable) (list []*AVLTreeNode) {
	if a == nil {
		return
	}

	if min.Compare(min, a.value) > 0 { // 节点值小于min，只往右找
		list = append(list, a.right.scope(min, max)...)
	} else if max.Compare(max, a.value) < 0 { // 节点值大于max，只往左找
		list = append(list, a.left.scope(min, max)...)
	} else {
		list = append(list, a.right.scope(min, max)...)
		list = append(list, a)
		list = append(list, a.left.scope(min, max)...)
	}
	return
}
