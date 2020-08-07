package linkedlists

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	header := head
	_ = header
	minList := &ListNode{
		Val:  -1,
		Next: nil,
	}
	maxList := &ListNode{
		Val:  -1,
		Next: nil,
	}

	minHead := minList
	maxHead := maxList
	for head != nil {
		if head.Val < x {
			minHead.Next = head
			head = head.Next
			minHead = minHead.Next
		} else {
			maxHead.Next = head
			head = head.Next
			maxHead = maxHead.Next
		}
	}
	maxHead.Next = nil
	minHead.Next = maxList.Next
	return minList.Next
}
