package linkedlists

func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	first.Next = swapPairsRecursion(second.Next)
	second.Next = first
	return second
}

func swapPairsIteration(head *ListNode) *ListNode {
	dump := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dump

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		pre.Next = second
		first.Next = second.Next
		second.Next = first

		pre = first
		head = first.Next
	}
	return dump.Next
}
