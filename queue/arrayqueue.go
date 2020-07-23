package queue

type ArrayQueue struct {
	items    []int
	length   uint
	capacity uint
}

func NewArrayQueue(capacity uint) *ArrayQueue {
	return &ArrayQueue{
		items:    make([]int, 0, capacity),
		length:   0,
		capacity: capacity,
	}
}

func (q *ArrayQueue) Enqueue(v int) bool {
	if q.length == q.capacity {
		return false
	}
	q.items = append(q.items, v)
	q.length++
	return true
}

func (q *ArrayQueue) Dequeue() (int, bool) {
	if q.length == 0 {
		return -1, false
	}
	result := q.items[0]
	q.items = q.items[1:]
	q.length--
	return result, true
}
