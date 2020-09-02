package heap

import (
	"container/heap"
)

type item struct {
	value    int
	priority int
	index    int // 记录对象在堆中的位置
}

type priorityQueue []item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq priorityQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }

func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(item)) }

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	res := old[n-1]
	*pq = old[:n-1]
	return res
}

// 修改优先队列中某个对象的优先级
func (pq *priorityQueue) update(item *item, value, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func topKFrequent(nums []int, k int) []int {
	dir := make(map[int]int)
	for _, num := range nums {
		dir[num]++
	}
	pq := new(priorityQueue)

	for k, v := range dir {
		*pq = append(*pq, item{value: k, priority: v})
	}

	heap.Init(pq)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(pq).(item).value
	}
	return res
}
