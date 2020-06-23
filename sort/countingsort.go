package sort

func CountingSort(list []int) []int {
	length := len(list)
	if length < 2 {
		return list
	}

	max := getMax(list, length)

	countArray := make([]int, max+1)

	for i := 0; i < length; i++ {
		// 计算每个元素的个数
		countArray[list[i]]++
	}

	for i := 1; i <= max; i++ {
		// 依次累加
		countArray[i] = countArray[i-1] + countArray[i]
	}

	// 存放有序的结果
	result := make([]int, length)

	for i := length - 1; i >= 0; i-- {
		result[countArray[list[i]]-1] = list[i]
		countArray[list[i]]--
	}

	copy(list, result)

	return list
}
