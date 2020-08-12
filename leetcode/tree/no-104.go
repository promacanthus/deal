package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	for len(queue) > 0 {
		size := len(queue) // 获取当前层队列的大小，控制深度
		for size > 0 {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			size--
		}
		depth++
	}
	return depth
}
