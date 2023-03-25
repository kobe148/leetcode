package _16

func isBalanced(root *TreeNode) bool {
	var maxDepth func(*TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxDepth(node.Left)
		right := maxDepth(node.Right)
		if left == -1 || right == -1 {
			return -1
		}

		if abs(left-right) > 1 {
			return -1
		}

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}

	return maxDepth(root) >= 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
