package queue

import "container/list"

type ListQueue struct {
	items    *list.List
	capacity uint
}

func NewListQueue(capacity uint) *ListQueue {
	return &ListQueue{items: list.New(), capacity: capacity}
}

func (q *ListQueue) Enqueue(v int) bool {
	if uint(q.items.Len()) == q.capacity {
		return false
	}
	q.items.PushBack(v)
	return true
}

func (q *ListQueue) Dequeue() (int, bool) {
	if q.items.Len() != 0 {
		v := q.items.Front()
		return v.Value.(int), true
	}
	return -1, false
}
