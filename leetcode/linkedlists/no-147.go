package linkedlists

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dump := &ListNode{}
	dump.Next = head

	insert := head.Next
	// 把已经有序的链的末尾断开
	// 相当于是两条链
	// 读取后一条链表头并在前一条中寻找合适位置插入
	head.Next = nil
	for insert != nil {
		prev := dump
		next := insert.Next
		// 寻找插入点
		for prev.Next != nil && prev.Next.Val <= insert.Val {
			prev = prev.Next
		}
		insert.Next = prev.Next
		prev.Next = insert
		insert = next
	}
	return dump.Next
}
