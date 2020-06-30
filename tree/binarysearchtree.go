package tree

type Node struct {
	data        int
	left, right *Node
	flag        bool
}

func NewNode(value int) *Node {
	return &Node{data: value, left: nil, right: nil, flag: false}
}

type BinarySearchTree struct {
	root  *Node
	count uint
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{root: nil, count: 0}
}

func (b *BinarySearchTree) Search(value int) bool {
	if b.count == 0 {
		return false
	}

	if search(b.root, value) != nil {
		return true
	}
	return false
}

func search(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if root.data == value {
		return root
	} else {
		if root.data < value {
			search(root.left, value)
		} else {
			search(root.right, value)
		}
	}
	return nil
}

func (b *BinarySearchTree) Insert(value int) {
	if value == b.root.data {
		return
	}
	findPosition(b.root, value).left = NewNode(value)
	b.count++
}

func findPosition(root *Node, value int) *Node {
	if value > root.data {
		if value < root.right.data {
			if root.right.left == nil {
				return root.right
			} else {
				findPosition(root.right.left, value)
			}
		} else if value > root.right.data {
			findPosition(root.right.right, value)
		}
		return root.right
	} else {
		if value < root.left.data {
			if root.left.left == nil {
				return root.left
			} else {
				findPosition(root.left.left, value)
			}
		} else {
			if value > root.left.data {
				findPosition(root.left.right, value)
			}
		}
		return root.left
	}
}

func (b *BinarySearchTree) Delete(value int) bool {
	if b.count == 0 {
		return false
	}
	search(b.root, value).flag = true
	return true
}

func (b *BinarySearchTree) Max() *Node {
	if b.count == 0 {
		return nil
	}

	child := b.root.right
	for child.right != nil {
		child = child.right
	}

	return child
}

func (b *BinarySearchTree) Min() *Node {
	if b.count == 0 {
		return nil
	}

	child := b.root.left
	for child.left != nil {
		child = child.left
	}

	return child
}

func (b *BinarySearchTree) PreOrder() []int {
	if b.count == 0 {
		return nil
	}
	res := make([]int, 0, b.count)
	preOrder(b.root, &res)
	return res
}

func preOrder(root *Node, res *[]int) {
	*res = append(*res, root.data)
	if root.left != nil {
		preOrder(root.left, res)
	}
	if root.right != nil {
		preOrder(root.right, res)
	}
}

func (b *BinarySearchTree) InOrder() []int {
	if b.count == 0 {
		return nil
	}
	res := make([]int, 0)
	inOrder(b.root, &res)
	return res
}

func inOrder(root *Node, res *[]int) {
	if root.left != nil {
		inOrder(root.left, res)
	}
	*res = append(*res, root.data)
	if root.right != nil {
		inOrder(root.right, res)
	}
}

func (b *BinarySearchTree) PostOrder() []int {
	if b.count == 0 {
		return nil
	}

	res := make([]int, 0)
	postOrder(b.root, &res)
	return res
}

func postOrder(root *Node, res *[]int) {
	if root.left != nil {
		postOrder(root.left, res)
	}

	if root.right != nil {
		postOrder(root.right, res)
	}
	*res = append(*res, root.data)
}
