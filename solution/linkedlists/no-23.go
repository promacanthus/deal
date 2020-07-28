package linkedlists

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
