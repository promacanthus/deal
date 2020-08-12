package queue

import (
	"testing"
)

func TestDoubleEndQueue(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		obj := Constructor(3)
		obj.InsertLast(1)
		obj.InsertLast(2)
		obj.InsertFront(3)
		obj.InsertFront(4)
		obj.GetRear()
		obj.IsFull()
		obj.DeleteLast()
		obj.InsertFront(4)
		obj.GetFront()
	})
}
