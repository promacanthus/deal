package linkedlists

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)
func hasCycleHashMap(head *ListNode) bool {
	res := make(map[*ListNode]bool)
	for head != nil {
		if res[head] {
			return true
		} else {
			res[head] = true
		}
		head = head.Next
	}
	return false
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	walker := head
	runner := head.Next
	for walker != runner {
		if runner == nil || runner.Next == nil {
			return false
		}
		walker = walker.Next
		runner = runner.Next.Next
	}
	return true
}