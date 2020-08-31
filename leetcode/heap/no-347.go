package heap

import "sort"

type priority struct {
	key int
	val int
}

type priorityQueue []priority

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].val > pq[j].val }

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(priority))
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

func topKFrequent(nums []int, k int) []int {
	if nums == nil || k < 1 {
		return nil
	}

	n := len(nums)
	if n < 1 {
		return nums
	}

	dir := make(map[int]int)
	for _, val := range nums {
		dir[val]++
	}

	pq := new(priorityQueue)
	for k, v := range dir {
		*pq = append(*pq, priority{key: k, val: v})
	}

	// heap.Init(pq)

	sort.Sort(pq)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = (*pq)[i].key
	}
	return res
}
