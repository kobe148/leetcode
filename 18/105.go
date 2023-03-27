package _18

// preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点
// 前序：根左右， 中序：左根右
func buildTree(preorder []int, inorder []int) *TreeNode {
	rootIndex := 0
	idxMap := make(map[int]int)
	for i := range inorder {
		idxMap[inorder[i]] = i
	}

	var dfs func(int, int) *TreeNode
	dfs = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		rootVal := preorder[rootIndex]
		rootIndex++
		root := &TreeNode{Val: rootVal}

		mid := idxMap[rootVal]
		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)

		return root
	}

	return dfs(0, len(inorder)-1)
}
