package avl

func formatValue(uid int32, exp uint32) (value uint64) {
	value = uint64(exp)<<32 | uint64(uid)
	return
}

func parseValue(value uint64) (uid int32) {
	uid = int32(value)
	return
}

//Max 获取2数中的较大值
func maxint32(x, y int32) int32 {
	if x < y {
		return y
	}
	return x
}

func maxint(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Comparable interface {
	Less(i, j interface{}) bool   // i < j return true
	Compare(i, j interface{}) int // i > j return 1, i == j return 0, i < j return -1
}
