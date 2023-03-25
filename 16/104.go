package _16

func maxDepth(root *TreeNode) int {
	var depth int

	var preOrder func(node *TreeNode, currDepth int)
	preOrder = func(node *TreeNode, currDepth int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if currDepth > depth {
				depth = currDepth
			}
		}
		preOrder(node.Left, currDepth+1)
		preOrder(node.Right, currDepth+1)
	}

	preOrder(root, 1)

	return depth
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 左右根
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
