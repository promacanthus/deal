package stack

type MinStack struct {
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{min: make([]int, 0), stack: make([]int, 0)}
}

func (s MinStack) Push(x int) {
	min := s.GetMin()
	if x < min {
		s.min = append(s.min, x)
	} else {
		s.min = append(s.min, min)
	}
	s.stack = append(s.stack, x)
}

func (s MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s MinStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

func (s MinStack) GetMin() int {
	if len(s.min) == 0 {
		return 1<<31 - 1
	}
	return s.min[len(s.min)-1]
}
