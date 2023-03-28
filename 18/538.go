package _18

// 反着的中序遍历
func convertBST(root *TreeNode) *TreeNode {
	currSum := 0

	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		currSum += node.Val
		node.Val = currSum
		dfs(node.Left)
	}

	dfs(root)
	return root
}
