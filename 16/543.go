package _16

func diameterOfBinaryTree(root *TreeNode) int {
	// 类似：求二叉树的最大深度, 有可能最大直径不会经过root节点
	ans := 0

	// 后序遍历
	var maxDepth func(*TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := maxDepth(node.Left)
		right := maxDepth(node.Right)
		if left+right > ans {
			ans = left + right
		}

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
	maxDepth(root)
	return ans
}
