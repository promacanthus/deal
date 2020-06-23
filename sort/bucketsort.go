package sort

func BucketSort(list []int) []int {
	length := len(list)
	if length < 2 {
		return list
	}

	max := getMax(list, length)

	buckets := make([][]int, length)
	index := 0

	for i := 0; i < length; i++ {
		index = list[i] * (length - 1) / max
		buckets[index] = append(buckets[index], list[i])
	}

	tmpPos := 0

	for i := 0; i < length; i++ {
		bucketLen := len(buckets)
		if bucketLen > 0 {
			QuickSort(buckets[i])
			copy(list[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}

	return list
}

func getMax(list []int, length int) int {
	max := list[0]

	for i := 0; i < length; i++ {
		if max < list[i] {
			max = list[i]
		}
	}

	return max
}
