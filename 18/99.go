package _18

// 中序遍历
// 另外一种解法：把树转为数组处理
func recoverTree(root *TreeNode) {
	var prev, x, y *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		if prev != nil && node.Val < prev.Val {
			// 第二个不符合条件的节点
			y = node
			// 第一个不符合条件的节点
			if x == nil {
				x = prev
			} else {
				return
			}
		}
		prev = node
		dfs(node.Right)
	}

	dfs(root)

	x.Val, y.Val = y.Val, x.Val
}
