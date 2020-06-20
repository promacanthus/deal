package linkedlist

type SingleNode struct {
	Value interface{}
	next  *SingleNode
}

func NewSingleNode(value interface{}) *SingleNode {
	return &SingleNode{Value: value, next: nil}
}

type DoubleNode struct {
	list      *DoubleLinkedList
	Value     interface{}
	pre, next *DoubleNode
}

func NewDoubleNode(value interface{}) *DoubleNode {
	return &DoubleNode{Value: value, pre: nil, next: nil}
}

func (d *DoubleNode) Next() *DoubleNode {
	if p := d.next; d.list != nil && d.list.head != p {
		return p
	}
	return nil
}

func (d *DoubleNode) Prev() *DoubleNode {
	if p := d.pre; d.list != nil && d.list.head != p {
		return p
	}
	return nil
}
