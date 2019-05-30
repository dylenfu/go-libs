package bubble

/*
冒泡排序
原理:
对给定的数组进行多次遍历，每次均比较相邻的两个数，
如果前一个比后一个大，则交换这两个数。
经过第一次遍历之后，最大的数就在最右侧了；
第二次遍历之后，第二大的数就在右数第二个位置了；以此类推。
*/

func BubbleSort(src []int) {
	if len(src) <= 1 {
		return
	}
	for i := 0; i < len(src) - 1; i++ {
		for j := i+1; j < len(src); j ++ {
			if src[i] > src[j] {
				src[i], src[j] = src[j], src[i]
			}
		}
	}
}
