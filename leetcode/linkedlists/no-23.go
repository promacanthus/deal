package linkedlists

import "container/heap"

// solution 1
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n < 1 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	return merge(lists, nil, 0, n)
}

func merge(l []*ListNode, ans *ListNode, i, n int) *ListNode {
	if i == n {
		return ans
	}

	if ans == nil {
		ans = mergeTwo(l[0], l[1])
		return merge(l, ans, 2, n)
	} else {
		ans = mergeTwo(ans, l[i])
		return merge(l, ans, i+1, n)
	}
}

func mergeTwo(a, b *ListNode) *ListNode {
	soldier := &ListNode{Val: 0, Next: nil}
	p := soldier

	for a != nil && b != nil {
		if a.Val <= b.Val {
			p.Next = a
			a = a.Next
		} else {
			p.Next = b
			b = b.Next
		}
		p = p.Next
	}

	if a != nil {
		p.Next = a
	}

	if b != nil {
		p.Next = b
	}

	return soldier.Next
}

// solution 2
type heapMin []*ListNode

func (h heapMin) Len() int            { return len(h) }
func (h heapMin) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h heapMin) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *heapMin) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *heapMin) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	n := len(lists)
	if n < 1 {
		return lists[0]
	}

	dump := new(ListNode)
	minHeap := new(heapMin)
	for i := 0; i < n; i++ {
		if lists[i] != nil {
			heap.Push(minHeap, lists[i])
		}
	}

	head := dump
	for minHeap.Len() != 0 {
		cur := heap.Pop(minHeap).(*ListNode)
		head.Next = cur
		head = head.Next
		if cur.Next != nil {
			heap.Push(minHeap, cur.Next)
		}
	}
	return dump.Next
}

// solution 3
func mergeKLists3(lists []*ListNode) *ListNode {
	return mgk(lists, 0, len(lists)-1)
}

func mgk(lists []*ListNode, low, high int) *ListNode {
	if low == high {
		return lists[low]
	}
	mid := low + (high-low)/2
	return mg2(mgk(lists, low, mid), mgk(lists, mid+1, high))
}

func mg2(a *ListNode, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val <= b.Val {
		a.Next = mg2(a.Next, b)
		return a
	}
	b.Next = mg2(a, b.Next)
	return b
}
