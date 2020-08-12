package tree

// 迭代
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		tmp := cur.Left
		cur.Left = cur.Right
		cur.Right = tmp
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return root
}

// 递归
func invertTreeRecursion(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTreeRecursion(root.Left)
	invertTreeRecursion(root.Right)
	return root
}
