package avl

import "testing"

type user struct {
	uid int32
	exp uint64
}

func (u *user) Compare(u1, u2 interface{}) int {
	exp1 := u1.(*user).exp
	exp2 := u2.(*user).exp

	if exp1 > exp2 {
		return 1
	} else if exp1 < exp2 {
		return -1
	} else {
		return 0
	}
}

func (u *user) Less(u1, u2 interface{}) bool {
	exp1 := u1.(*user).exp
	exp2 := u2.(*user).exp

	return exp1 < exp2
}

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree()
	list := []int{72, 38, 2, 65, 42, 44, 39, 8}
	for i, value := range list {
		u := &user{uid: int32(10000 + i), exp: uint64(value)}
		tree.Insert(u)
	}
	for _, v := range tree.Scan(tree.head) {
		t.Log(v.value)
	}
}
