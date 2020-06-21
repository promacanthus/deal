package queue

import "container/list"

type ListQueue struct {
	items    *list.List
	capacity uint
}

func NewListQueue(capacity uint) *ListQueue {
	return &ListQueue{
		items:    list.New(),
		capacity: capacity}
}

func (l *ListQueue) Enqueue(n int) bool {
	if uint(l.items.Len()) == l.capacity {
		return false
	}
	l.items.PushBack(n)
	return true
}

func (l *ListQueue) Dequeue() (int, bool) {
	if l.items.Len() != 0 {
		element := l.items.Front()
		return element.Value.(int), true
	}
	return -1, false
}
