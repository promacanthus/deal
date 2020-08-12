package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSumIteration(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSumIteration(root.Left, sum-root.Val) || hasPathSumIteration(root.Right, sum-root.Val)
}

func hasPathSumBFS(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	nodeQueue := make([]*TreeNode, 0)
	sumQueue := make([]int, 0)
	nodeQueue = append(nodeQueue, root)
	sumQueue = append(sumQueue, root.Val)
	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		tmpSum := sumQueue[0]
		sumQueue = sumQueue[1:]
		if node.Left == nil && node.Right == nil {
			if tmpSum == sum {
				return true
			}
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			sumQueue = append(sumQueue, node.Left.Val+tmpSum)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			sumQueue = append(sumQueue, node.Right.Val+tmpSum)
		}
	}
	return false
}
