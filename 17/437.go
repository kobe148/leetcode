package _17

// DFS(前序遍历) + 前缀和
func pathSum437(root *TreeNode, targetSum int) int {
	var res, prefixSumMap = 0, make(map[int]int)
	// prefixSumMap前缀和集合
	prefixSumMap[0] = 1

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, currSum int) {
		if node == nil {
			return
		}
		currSum += node.Val
		// 在前缀和集合里面查找对应的前缀和出现的次数
		res += prefixSumMap[currSum-targetSum]
		prefixSumMap[currSum] += 1

		dfs(node.Left, currSum)
		dfs(node.Right, currSum)

		prefixSumMap[currSum] -= 1
	}

	dfs(root, 0)

	return res
}
