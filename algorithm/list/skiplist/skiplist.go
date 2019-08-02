package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	skiplist_rand_p    int32 = 2
	skiplist_max_level       = 4
)

// 使用score int类型数据排序
type Element struct {
	score   int
	value   interface{}
	forward []*Element
}

func newElement(score int, value interface{}, level int) *Element {
	return &Element{
		score:   score,
		value:   value,
		forward: make([]*Element, level),
	}
}

type SkipList struct {
	header *Element
	len    int
	level  int
}

func New() *SkipList {
	return &SkipList{
		header: &Element{forward: make([]*Element, skiplist_max_level)},
	}
}

// Search search
func (list *SkipList) Search(score int) (*Element, bool) {
	e := list.header

	// 找到距离目标分数最近但并不相等的节点
	for i := list.level - 1; i >= 0; i-- {
		for e.forward[i] != nil && e.forward[i].score < score {
			e = e.forward[i]
		}
	}

	// 现在不知道节点e有几层指针，但可以肯定的是该节点最底层的指针最接近目标值
	e = e.forward[0]
	if e != nil && e.score == score {
		return e, true
	}
	return nil, false
}

// Insert insert the aiming k-v data
func (list *SkipList) Insert(score int, value interface{}) *Element {
	// update用来存储不同层级指向插入节点的节点指针
	update := make([]*Element, skiplist_max_level)

	// 这里需要注意,链表有几层，那么update也就有几层
	e := list.header
	for i := list.level - 1; i >= 0; i-- {
		for e.forward[i] != nil && e.forward[i].score < score {
			e = e.forward[i]
		}
		update[i] = e
	}

	// 如果已经存在则直接返回
	e = e.forward[0]
	if e != nil && e.score == score {
		return e
	}

	// 新插入的节点需要链接到几层指针是随机的，只要有一层相连就可以找到后续元素
	// 如果随机得到的level比链表当前层数更大，每次插入元素时，也最多只能在链表
	// 现有层数上加1层; 同时，update中新加的层级元素需要赋值为list.header
	// list的level是从0开始，而不是从1开始
	level := randLevel()
	if level > list.level {
		// level = list.level + 1
		// update[level-1] = list.header
		for i := list.level; i < level; i++ {
			update[i] = list.header
		}
		list.level = level
	}

	// 新建需要插入的节点p, 将p节点插入到单链表中
	p := newElement(score, value, list.level)
	for i := 0; i < list.level; i++ {
		q := update[i]
		p.forward[i] = q.forward[i]
		q.forward[i] = p
	}
	list.len++

	return p
}

// Delete remove element from skip list, return nil if element not exist
func (list *SkipList) Delete(score int) *Element {
	// 存储每层需要断开的节点指针
	update := make([]*Element, skiplist_max_level)

	e := list.header
	for i := list.level - 1; i >= 0; i-- {
		for e.forward[i] != nil && e.forward[i].score < score {
			e = e.forward[i]
		}
		update[i] = e
	}

	p := e.forward[0]
	if p != nil && p.score == score {
		for i := 0; i < list.level; i++ {
			q := update[i]
			q.forward[i] = p.forward[i]
			p.forward[i] = nil
		}
		return p
	} else {
		return nil
	}
}

func (list *SkipList) PrintList() {
	for i := list.level - 1; i >= 0; i-- {
		e := list.header
		num := 0
		for {
			// fmt.Println(fmt.Sprintf("%d", e.score))
			num++
			if e.forward[i] != nil {
				e = e.forward[i]
			} else {
				break
			}
		}
		fmt.Println(fmt.Sprintf("level %d num %d", i, num))
		fmt.Println("---------------------------------------------")
	}
}

func randLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())
	for (rand.Int31()&0xffff)%skiplist_rand_p == 0 {
		level++
	}
	if level > skiplist_max_level {
		level = skiplist_max_level
	}
	fmt.Println(fmt.Sprintf("---------hahahhahhahah %d", level))
	return level
}
