package search

func BinarySearch(list []int, value int) bool {
	if len(list) < 1 {
		return false
	}
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if value == list[mid] {
			return true
		} else {
			if list[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}

func BinarySearchRecursive(list []int, value int) bool {
	if len(list) < 1 {
		return false
	}
	return bsr(list, 0, len(list)-1, value)
}

func bsr(list []int, low, high, value int) bool {
	if low >= high {
		return false
	}
	mid := low + ((high - low) >> 1)
	if list[mid] == value {
		return true
	} else {
		if list[mid] < value {
			return bsr(list, mid+1, high, value)
		} else {
			return bsr(list, low, mid-1, value)
		}
	}
}

// 查找第一个等于给定值的索引
func BinarySearchFirst(list []int, value int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if list[mid] < value {
			low = mid + 1
		} else if list[mid] > value {
			high = mid - 1
		} else {
			if mid == 0 || list[mid-1] != value {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// 查找最后一个等于给定值的索引
func BinarySearchLast(list []int, value int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if list[mid] < value {
			low = mid + 1
		} else if list[mid] > value {
			high = mid - 1
		} else {
			if mid == len(list)-1 || list[mid+1] != value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

// 查找第一个大于等于给定值的索引
func BinarySearchGE(list []int, value int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if list[mid] >= value {
			if mid == 0 || list[mid-1] < value {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的索引
func BinarySearchLE(list []int, value int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + ((high - low) >> 1)
		if list[mid] > value {
			high = mid - 1
		} else {
			if mid == len(list)-1 || list[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
