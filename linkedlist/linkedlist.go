package linkedlist

import (
	"errors"
	"fmt"
)

/*
	1.理解指针
    2.警惕指针丢失和内存泄漏
	3.利用哨兵简化实现难度
	4.注意边界条件处理
	5.举例画图辅助思考
*/

type Element struct {
	value int
	next  *Element
}

type LinkedList struct {
	head   *Element
	length uint
}

func NewElement(v int) *Element {
	return &Element{
		value: v,
		next:  nil,
	}
}

func (e *Element) GetNext() *Element {
	return e.next
}

func (e *Element) GetValue() int {
	return e.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   NewElement(0),
		length: 0,
	}
}

// 在指定节点后插入
func (l *LinkedList) InsertAfter(p *Element, v int) bool {
	if p == nil {
		return false
	}
	element := NewElement(v)
	oldNext := p.next
	p.next = element
	element.next = oldNext
	l.length++
	return true
}

// 在指定节点前插入
func (l *LinkedList) InsertBefore(p *Element, v int) bool {
	// p不为空，不为头指针
	if p == nil || p == l.head {
		return false
	}

	cur := l.head.next
	pre := l.head

	// 链表存取其他节点（非头指针）
	if cur == nil {
		return false
	}

	// 遍历P的位置
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	element := NewElement(v)

	pre.next = element
	element.next = cur
	l.length++
	return true
}

// 在链表头插入
func (l *LinkedList) InsertAtHeader(v int) bool {
	return l.InsertAfter(l.head, v)
}

// 在链表尾部插入
func (l *LinkedList) InsertAtTail(v int) bool {
	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	return l.InsertAfter(cur, v)
}

// 通过索引查找节点
func (l *LinkedList) FindByIndex(n uint) *Element {
	if n >= l.length {
		return nil
	}

	cur := l.head
	for i := uint(0); i <= n; i++ {
		cur = cur.next
	}

	return cur
}

// 删除节点
func (l *LinkedList) DeleteNode(p *Element) bool {
	if p == nil || l.length == 0 {
		return false
	}

	cur := l.head.next
	pre := l.head

	if cur == nil {
		return false
	}

	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	pre.next = p.next
	p = nil
	l.length--
	return true
}

func (l *LinkedList) Print() {
	format := ""
	cur := l.head.next
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

// 反转链表
func (l *LinkedList) Reverse() {
	cur := l.head.next.next
	pre := l.head.next

	for cur.next != nil {
		pos := cur.next
		pre.next = l.head
		cur.next.next = cur
		cur = pos
	}
}

// 判断是否存在环
func (l *LinkedList) ExistCycle() bool {
	return false
}

// 合并两个有序链表
func (l *LinkedList) Combine(ll *LinkedList) error {
	if !l.ordered() || !ll.ordered() {
		return errors.New("linked list is not ordered")
	}

	pre1 := l.head
	cur1 := l.head.next
	cur2 := ll.head.next

	for {
		if cur1.value >= cur2.value {
			tmp1 := cur1.next
			tmp2 := cur2.next
			cur1.next = cur2
			cur2.next = tmp1
			cur1 = tmp1
			cur2 = tmp2
		} else {
			tmp2 := cur2.next
			pre1.next = cur2
			cur2.next = cur1
			cur2 = tmp2
		}
		cur1 = cur1.next
		cur2 = cur2.next
	}
	return nil
}

func (l *LinkedList) ordered() bool {
	cur := l.head.next.next
	pre := l.head.next

	if cur.value >= pre.value {
		for cur.next != nil {
			if cur.value < pre.value {
				return false
			}
		}
	} else {
		for cur.next != nil {
			if cur.value >= pre.value {
				return false
			}
		}
	}
	return true
}

// 删除倒数低n个

// 查找中间节点
