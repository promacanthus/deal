package tree

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	res := make([]int, 0)
	inOrder(root, &res)

	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}

	return true
}

func inOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, res)
	*res = append(*res, node.Val)
	inOrder(node.Right, res)
}

func isValidBSTRecursion(root *TreeNode) bool {
	return recursion(root, math.MinInt64, math.MaxInt64)
}

func recursion(node *TreeNode, low, up int) bool {
	if node == nil {
		return true
	}
	if node.Val <= low || node.Val >= up {
		return false
	}
	return recursion(node.Left, low, node.Val) && recursion(node.Right, node.Val, up)
}
