type AVLTree struct {
	root *AVLTreeNode
}

// avl
type AVLTreeNode struct {
	height int32        //节点高
	value  uint64       //节点值
	left   *AVLTreeNode //节点左儿子
	right  *AVLTreeNode //节点右儿子
}

func NewNode(uid int32, exp uint32,left,right *AVLTreeNode) (entry *AVLTreeNode) {
	entry = &AVLTreeNode{value: formatValue(uid, exp), left:left, right:right}
	return
}

func NewTree() *AVLTree {
	return &AVLTree{}
}

func (a *AVLTreeNode) getHeight() int32 {
	if a == nil {
		return 0
	} else {
		return a.height
	}
}

//updateHeight 更新节点高度
func (a *AVLTreeNode) updateHeight() {
	a.height = max(a.left.getHeight(), a.right.getHeight()) + 1
}

//balanceFactor 获取节点左右子树高度差绝对值
//将二叉树上节点的左子树高度减去右子树高度取绝对值(Balance Factor)
func (a *AVLTreeNode) balanceFactor() int32 {
	var lh, rh int32
	if a == nil {
		return 0
	}
	lh = a.left.getHeight()
	rh = a.right.getHeight()
	if lh > rh {
		return lh - rh
	} else {
		return rh - lh
	}
}

//FindMax 查找最大节点
func (t *AVLTree) FindMax(a *AVLTreeNode) *AVLTreeNode {
	if a == nil {
		return nil
	}
	if a.right != nil {
		return t.FindMax(a.right)
	}
	return a
}

//FindMin 查找最小值
func (t *AVLTree) FindMin(a *AVLTreeNode) *AVLTreeNode {
	if a == nil {
		return nil
	}
	if a.left != nil {
		return t.FindMin(a.left)
	}
	return a
}

/*
		   x             y
		  / \           / \
		t1   y  ->     x   t3
            / \       / \
           t2 t3     t1 t2

         pic-1         pic-2

如图中pic-1所示，如果再继续在右子树插入右孩子，会导致AVL失衡
为了处理这种情况，需要将x节点左旋，降低右子树高度。然后再插入右孩子

LeftRotation 左旋，返回该子树新的根节点(即pic-2中y)
*/
func (t *AVLTree) LeftRotation(x *AVLTreeNode) *AVLTreeNode {
	y := x.right
	x.right = y.left
	y.left = x
	x.updateHeight()
	y.updateHeight()
	return y
}

/*
			x                  y
           / \               /  \
          y   t3    ->      t1   x
         / \                    /  \
        t1  t2                 t2  t3

          pic-1                pic-2

如图中pic-1所示，如果要继续往左子树插入左孩子，会导致AVL失衡
为了处理这种情况，需要将y右旋，降低左子树高度，然后再插入左孩子

RightRotation 右旋，返回该子树新的根节点(即pic-2中x)
*/
func (t *AVLTree) RightRotation(x *AVLTreeNode) *AVLTreeNode {
	y := x.left
	x.left = y.right
	y.right = x
	x.updateHeight()
	x.updateHeight()
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

LeftRightRotation 先左后右，具体解决往左子树插右孩子的具体情况
*/
func (t *AVLTree) LeftRightRotation(x *AVLTreeNode) *AVLTreeNode {
	x.left = t.LeftRotation(x.left)
	return t.RightRotation(x)
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

RightLeftRotation 先右后左，解决往右子树插入左节点的场景
*/
func (t *AVLTree) RightLeftRotation(a *AVLTreeNode) *AVLTreeNode {
	a.right = t.RightRotation(a.right)
	return t.LeftRotation(a)
}

//InsertAVL avl树插入操作
//递归
// 1个节点的高度初始值为1
//返回插入后的根节点
func (t *AVLTree) InsertAVL(tree *AVLTreeNode, v uint64) *AVLTreeNode {
	if tree == nil { //空树，插入第一个节点
		tree := new(AVLTreeNode)
		tree.value = v
	} else if v < tree.value { //待插入的值比单前节点值小，往左子树插入
		tree.left = t.InsertAVL(tree.left, v)
		if tree.left.getHeight() - tree.right.getHeight() == 2 { //插入新节点后avl树失衡
			if v < tree.left.value { //情况1：v插入到左子树的左孩子节点，只需要进行一次右旋转
				tree = t.RightRotation(tree)
			} else if v > tree.left.value { //情况2：v插入到左子树的右孩子节点，需要先左旋转再右旋转
				tree = t.LeftRightRotation(tree)
			}
		}
	} else if v > tree.value {
		tree.right = t.InsertAVL(tree.right, v)
		if tree.right.getHeight() - tree.left.getHeight() == 2 { //在右子树插入新节点后avl树失衡
			if v > tree.right.value { //情况4：v插入到右子树的右孩子节点，只需要进行一次左旋转
				tree = t.LeftRotation(tree)
			} else if v < tree.right.value { //情况3：v插入到右子树的左孩子节点，需要先右旋转再左旋转
				tree = t.RightLeftRotation(tree)
			}
		}
	}
	tree.updateHeight()
	return tree
}

//DeleteAVL 删除avl树中值为v的节点
//维护二叉排序树规则
//平衡avl二叉树, 返回新的根节点
func (t *AVLTree) DeleteAVL(tree *AVLTreeNode, v uint64) *AVLTreeNode {
	if tree == nil {
		return nil
	}
	if tree.value == v { //tree为待删除的节点
		//删除后维护成新的二叉排序树
		if tree.left != nil && tree.right != nil { //若待删除节点同时存在左右子树
			if tree.left.getHeight() > tree.right.getHeight() { //左子树高于右子树，则取左子树最大值取代被删除节点
				tmp := t.FindMax(tree.left)
				tree.value = tmp.value
				tree.left = t.DeleteAVL(tree.left, tmp.value)
			} else { //右子树高于左子树，则取右子树最小值取代被删除节点
				tmp := t.FindMin(tree.right)
				tree.value = tmp.value
				tree.right = t.DeleteAVL(tree.right, tmp.value)
			}
		} else {
			if tree.left != nil {
				tree = tree.left
			}
			if tree.right != nil {
				tree = tree.right
			}
		}
		return tree
	} else if tree.value < v { // 待删除节点比当前节点大
		tree.right = t.DeleteAVL(tree.right, v)
		if tree.left.getHeight() - tree.right.getHeight() == 2 { //删除节点后avl树失衡
			if v < tree.left.value { //情况1：左子树的左孩子节点过高，只需要进行一次右旋转
				tree = t.RightRotation(tree)
			} else if v > tree.left.value { //情况2：左子树的右孩子节点过高，需要先左旋转再右旋转
				tree = t.LeftRightRotation(tree)
			}
		}
	} else if tree.value > v { // 待删除节点比当前节点小
		tree.left = t.DeleteAVL(tree.left, v)
		if tree.right.getHeight() - tree.left.getHeight() == 2 { //在右子树插入新节点后avl树失衡
			if v > tree.right.value { //情况4：右子树的右孩子节点过高，只需要进行一次左旋转
				tree = t.LeftRotation(tree)
			} else if v < tree.right.value { //情况3：右子树的左孩子节点过高，需要先右旋转再左旋转
				tree = t.RightLeftRotation(tree)
			}
		}
	}
	return tree
}

