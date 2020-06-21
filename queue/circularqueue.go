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

func (c *CircularQueue) Enqueue(elem int) bool {
	if (c.tail+1)%c.capacity != c.head {
		c.items[c.tail] = elem
		c.tail = (c.tail + 1) % c.capacity
		return true
	}
	return false
}

func (c *CircularQueue) Dequeue() (int, bool) {
	if c.head != c.tail {
		elem := c.items[c.head]
		c.head = (c.head + 1) % c.capacity
		return elem, true
	}
	return -1, false
}
