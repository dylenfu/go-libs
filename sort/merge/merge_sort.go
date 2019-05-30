package merge

/*
  说明:
  以[3, 1, 2, 5, 6, 43, 4]为例
  第1步: 将数组拆分成[3, 1, 2], [5, 6, 43, 4]
  第2步: 将[3, 1, 2]拆分成[3], [1, 2]
  第3步: [3], 直接返回, [1, 2]进一步拆分成[1], [2], 并在merge中进行排序及合并
  第4步: [3], 合并后的[1, 2]在merge中合并:
    l, r = 0
    1, 2都比3小 所以1, 2 append 到result 同时r++
    [3]直接append到result
  第5步: [5, 6, 43, 4]重复2~4过程
  第6步: 排序并合并后的[1, 2, 3] & [4, 5, 6, 43]进行merge获得结果
*/

// 将数组递归拆解成最小单元,然后逐步合并成已排序数组
func MergeSort(src []int) []int {
	length := len(src)

	if length <= 1 {
		return src
	}

	mid := length / 2
	left := MergeSort(src[:mid])
	right := MergeSort(src[mid:])
	return merge(left, right)
}

// 合并两个数组到一个数组, 合并过程中排序
// 1.设定两个数组的起始index
// 2.排序时当找到一个min值是该index右移,同时将min值append到result
// 3.比较过程中剩下的数据直接append到result数组
func merge(left, right []int) (result []int) {
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
