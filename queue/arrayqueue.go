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

func (a *ArrayQueue) Enqueue(elem int) bool {
	if a.length == a.capacity {
		return false
	}
	a.items = append(a.items, elem)
	a.length++
	return true
}

func (a *ArrayQueue) Dequeue() (int, bool) {
	if a.length != 0 {
		result := a.items[0]
		a.items = a.items[1:]
		a.length--
		return result, true
	}

	return -1, false
}
