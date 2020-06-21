package stack

type ArrayStack struct {
	items    []int
	length   uint
	capacity uint
}

func NewStack(n uint) *ArrayStack {
	return &ArrayStack{
		items:    make([]int, 0, n),
		length:   0,
		capacity: n,
	}
}

func (s *ArrayStack) Push(n int) bool {
	if s.length == s.capacity {
		return false
	}
	s.items = append(s.items, n)
	s.length++
	return true
}

func (s *ArrayStack) Pop() (int, bool) {
	if s.length == 0 {
		return -1, false
	}
	result := s.items[s.length-1]
	s.items = s.items[:s.length-1]
	s.length--
	return result, true
}
