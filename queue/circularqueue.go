package queue

type CircularQueue struct {
	items                []int
	head, tail, capacity uint
}

func NewCircularQueue(capacity uint) *CircularQueue {
	return &CircularQueue{
		items:    make([]int, 0, capacity),
		head:     0,
		tail:     0,
		capacity: capacity,
	}
}

func (q *CircularQueue) Enqueue(v int) bool {
	if (q.tail+1)%q.capacity != q.head {
		q.items[q.tail] = v
		q.tail = (q.tail + 1) % q.capacity
		return true
	}
	return false
}

func (q *CircularQueue) Dequeue() (int, bool) {
	if q.head != q.tail {
		v := q.items[q.head]
		q.head = (q.head + 1) % q.capacity
		return v, true
	}
	return -1, false
}
