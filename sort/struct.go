package sort

import (
	"math/big"
	tsort "sort"
)

// implement sort interface functions len,swap,less

type Student struct {
	Name string
	Height *big.Int
}

type Students []*Student

func (s Students) Len() int{
	return len(s)
}

func (s Students) Less(i, j int) bool {
	si := s[i]
	sj := s[j]
	if si.Height.Cmp(sj.Height) < 0  {
		return true
	}

	return false
}

func (s Students) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func SimpleStructSort() {
	s1 := &Student{"s1", big.NewInt(20)}
	s2 := &Student{"s2", big.NewInt(19)}
	s3 := &Student{"s3", big.NewInt(20)}
	ss := Students{s1, s2, s3}
	tsort.Sort(ss)

	for _, v := range ss {
		println(v.Name, v.Height.String())
	}
}

