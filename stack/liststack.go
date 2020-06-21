package stack

import (
	"container/list"
)

type ListStack struct {
	items    *list.List
	length   uint
	capacity uint
}

func NewListStack(n uint) *ListStack {
	return &ListStack{
		items:    list.New(),
		length:   0,
		capacity: n,
	}
}

func (s *ListStack) Push(n int) bool {
	if s.length == s.capacity {
		return false
	}

	s.items.PushBack(n)
	s.length++
	return true
}

func (s *ListStack) Pop() (int, bool) {
	if s.length == 0 {
		return -1, false
	}

	result := s.items.Back()
	s.items.Remove(result)
	s.length--

	return result.Value.(int), true
}
