package linkedlists

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dump := &ListNode{}
	dump.Next = head

	insert := head.Next
	for insert != nil {
		pre := dump
		sorted := dump.Next
		for sorted.Val < insert.Val {
			sorted = sorted.Next
			pre = pre.Next
		}
		pre.Next = insert
		next := insert.Next
		insert.Next = sorted
		sorted.Next = next
		insert = next
	}
	return dump.Next
}
