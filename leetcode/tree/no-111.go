package tree

// 注意是叶子节点
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	minV := 1
	if root.Left != nil {
		minV = min(minDepth(root.Left), minV)
	}
	if root.Right != nil {
		minV = min(minDepth(root.Right), minV)
	}
	return minV + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	find := false
	depth := 0
	for len(queue) > 0 {
		size := len(queue)
		// 控制层数，在一个循环中入队的是在同一层
		for size > 0 {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left == nil && curNode.Right == nil {
				find = true
				break
			}
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			size--
		}
		depth++
		if find {
			break
		}
	}
	return depth
}
