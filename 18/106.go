package _18

func buildTree2(inorder []int, postorder []int) *TreeNode {
	rootIndex := len(postorder) - 1
	idxMap := make(map[int]int)
	for i := range inorder {
		idxMap[inorder[i]] = i
	}

	var dfs func(int, int) *TreeNode
	dfs = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := postorder[rootIndex]
		rootIndex--
		root := &TreeNode{Val: rootVal}

		mid := idxMap[rootVal]

		root.Right = dfs(mid+1, right)
		root.Left = dfs(left, mid-1)

		return root
	}

	return dfs(0, len(inorder)-1)
}
