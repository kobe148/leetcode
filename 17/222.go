package _17

// 后序遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countNodes(root.Left)
	right := countNodes(root.Right)

	return left + right + 1
}
