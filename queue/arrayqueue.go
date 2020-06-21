package queue

type SliceQueue struct {
	items    []int
	length   uint
	capacity uint
}

func NewSliceQueue(capacity uint) *SliceQueue {
	return &SliceQueue{
		items:    make([]int, 0, capacity),
		length:   0,
		capacity: capacity,
	}
}

func (s *SliceQueue) Enqueue(elem int) bool {
	if s.length == s.capacity {
		return false
	}
	s.items = append(s.items, elem)
	s.length++
	return true
}

func (s *SliceQueue) Dequeue() (int, bool) {
	if s.length != 0 {
		result := s.items[0]
		s.items = s.items[1:]
		s.length--
		return result, true
	}

	return -1, false
}
