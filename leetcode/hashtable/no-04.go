package hashtable

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	all := n1 + n2

	res := make([]int, 0)
	res = append(res, nums1...)
	res = append(res, nums2...)

	sort.Ints(res)
	if all%2 != 0 {
		index := all / 2
		return float64(res[index])
	} else {
		index := all / 2
		return float64(res[index-1]+res[index]) / 2
	}
}
