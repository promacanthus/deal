package linkedlist

import (
	"fmt"
)

type SingleLinkedList struct {
	head   *SingleNode
	length uint
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{head: NewSingleNode(nil), length: 0}
}

// 在指定节点后插入
func (s *SingleLinkedList) InsertAfter(p *SingleNode, v interface{}) bool {
	if p == nil {
		return false
	}
	node := NewSingleNode(v)
	oldNext := p.next
	p.next = node
	node.next = oldNext
	s.length++
	return true
}

// 在链表头插入
func (s *SingleLinkedList) First(v int) bool {
	return s.InsertAfter(s.head, v)
}

// 在指定节点前插入
func (s *SingleLinkedList) InsertBefore(p *SingleNode, v interface{}) bool {
	// p不为空，不为头指针
	if p == nil || p == s.head {
		return false
	}

	cur := s.head.next
	pre := s.head

	// 链表存取其他节点（非头指针）
	if cur == nil {
		return false
	}

	// 寻找p的位置
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	node := NewSingleNode(v)
	pre.next = node
	node.next = cur
	s.length++
	return true
}

// 在链表尾部插入
func (s *SingleLinkedList) Last(v int) bool {
	cur := s.head

	for cur.next != nil {
		cur = cur.next
	}

	return s.InsertAfter(cur, v)
}

// 通过索引查找节点
// index = 0 表示 head
func (s *SingleLinkedList) FindByIndex(index uint) *SingleNode {
	if index > s.length || s.head.next == nil {
		return nil
	}

	cur := s.head
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}

	return cur
}

// 删除节点
func (s *SingleLinkedList) Remove(p *SingleNode) bool {
	if p == nil || s.length == 0 {
		return false
	}

	cur := s.head.next
	pre := s.head

	// 寻找p的位置
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	pre.next = cur.next
	p = nil
	s.length--
	return true
}

//Print 打印单链表
func (s *SingleLinkedList) Print() {
	format := ""
	cur := s.head.next
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.Value)
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

// 反转链表
func (s *SingleLinkedList) Reverse() {
	pre := s.head.next
	cur := s.head.next.next

	s.head.next = nil // 把head与链断开
	pre.next = nil    // 原链首变链尾

	for cur.next != nil {
		next := cur.next // 保存下一个节点
		cur.next = pre   // 指针反转
		pre = cur        // 保存前一个节点
		cur = next       // 指针递增
	}
	cur.next = pre
	s.head.next = cur
}

// 判断是否存在环
func (s *SingleLinkedList) ExistCycle() bool {
	if s.head.next == nil {
		return false
	}

	walker := s.head
	runner := s.head

	for runner.next != nil || runner.next.next != nil {
		walker = walker.next
		runner = runner.next.next
		if walker == runner {
			return true
		}
	}
	return false
}

// 合并两个有序链表
func (s *SingleLinkedList) Merge(sll *SingleLinkedList) *SingleLinkedList {
	if s.head.next == nil {
		return sll
	}

	if sll.head.next == nil {
		return s
	}

	result := NewSingleLinkedList()
	cur := result.head
	cur1 := s.head
	cur2 := sll.head
	for cur1.next != nil && cur2.next != nil {
		// value需要可比
		if cur1.next.Value < cur2.next.Value {
			cur.next = cur1.next
			cur1 = cur1.next
		} else {
			cur.next = cur2.next
			cur2 = cur2.next
		}
		cur = cur.next
	}
	return result
}

// 删除倒数第n个
// 即正数第length-n+1
func (s *SingleLinkedList) RemoveLastN(n uint) bool {
	return s.Remove(s.FindByIndex(s.length - n + 1))
}

// 查找中间节点
func (s *SingleLinkedList) Middle() *SingleNode {
	switch s.length {
	case 1, 2:
		return s.head.next
	default:
		fast := s.head
		slow := s.head
		if s.length%2 != 0 {
			// 奇数
			for fast.next.next != nil {
				fast = fast.next.next
				slow = slow.next
			}
			return slow.next
		}
		// 偶数
		for fast.next != nil {
			fast = fast.next.next
			slow = slow.next
		}
		return slow
	}
}
