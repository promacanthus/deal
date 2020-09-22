package stack

import "testing"

func Test(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.GetMin()
	obj.Pop()
	obj.Top()
	obj.GetMin()
}
