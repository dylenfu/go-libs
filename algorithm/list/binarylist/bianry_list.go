package binarylist

import (
	"sync"
)

type list struct {
	data []uint64
	mtx *sync.Mutex
}

func NewList() *list{
	return &list{data: make([]uint64, 32), mtx: new(sync.Mutex)}
}

func (l *list) Insert(value uint64) bool {

}

func (l *list) Delete(value uint64) bool {

}

func (l *list) Range(start, end int) []uint64 {

}

// find return index and judgement state
func (l *list) find(value uint64, start, end int) (int, bool) {
	len := len(l.data)
	x := len / 2

}
