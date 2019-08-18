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

func (u *user) Min(u1, u2 Comparable) Comparable {
	if u.Compare(u1, u2) < 0 {
		return u1
	}
	return u2
}

func (u *user) Max(u1, u2 Comparable) Comparable {
	if u.Compare(u1, u2) > 0 {
		return u1
	}
	return u2
}

func (u *user) InScore(min, max, value Comparable) bool {
	if value.Compare(value, min) < 0 {
		return false
	} else if value.Compare(value, max) > 0 {
		return false
	} else {
		return true
	}
}

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree()
	list := []int{72, 38, 2, 65, 42, 44, 39, 8}
	for i, value := range list {
		u := &user{uid: int32(10000 + i), exp: uint64(value)}
		tree.Insert(u)
	}
	t.Log("-------scan----------")
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}
	t.Log("--------get range------")
	min := &user{exp: uint64(26)}
	max := &user{exp: uint64(45)}
	for _, v := range tree.Range(min, max) {
		t.Log(v.value)
	}
	t.Log("--------delete not exist 26---------")
	tree.Delete(min)
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}

	t.Log("--------delete 38---------")
	tree.Delete(&user{exp: uint64(38)})
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}

	t.Log("--------delete 72---------")
	tree.Delete(&user{exp: uint64(72)})
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}

	t.Log("--------delete 65---------")
	tree.Delete(&user{exp: uint64(65)})
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}

	t.Log("--------delete 44---------")
	tree.Delete(&user{exp: uint64(44)})
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}

	t.Log("--------delete 42---------")
	tree.Delete(&user{exp: uint64(42)})
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}

	t.Log("--------delete 39---------")
	tree.Delete(&user{exp: uint64(39)})
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}

	t.Log("--------delete 8---------")
	tree.Delete(&user{exp: uint64(8)})
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}

	t.Log("--------delete 2---------")
	tree.Delete(&user{exp: uint64(2)})
	for _, v := range tree.Scan() {
		t.Log(v.value)
	}
}
