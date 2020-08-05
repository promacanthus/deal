package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := make([][]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		res := make([]int, 0)
		for size > 0 {
			curNode := queue[0]
			queue = queue[1:]
			res = append(res, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			size--
		}
		ans = append(ans, res)
	}
	return ans
}
