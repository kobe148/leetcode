package _17

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val == 0
	}

	left := hasPathSum(root.Left, targetSum-root.Val)
	if left {
		return true
	}
	right := hasPathSum(root.Right, targetSum-root.Val)

	return left || right
}

func hasPathSum3(root *TreeNode, targetSum int) bool {
	var dfs func(node *TreeNode, parentNodeTarget int, res *[]int)
	dfs = func(node *TreeNode, parentNodeTarget int, res *[]int) {
		if node == nil {
			return
		}
		currNodeTarget := parentNodeTarget - node.Val
		if node.Left == nil && node.Right == nil {
			*res = append(*res, currNodeTarget)
		}

		dfs(node.Left, currNodeTarget, res)
		dfs(node.Right, currNodeTarget, res)
	}

	n := make([]int, 0)
	dfs(root, targetSum, &n)
	for i := 0; i < len(n); i++ {
		if n[i] == 0 {
			return true
		}
	}

	return false
}

func hasPathSum2(root *TreeNode, targetSum int) bool {
	resList := allPath(root)

	for _, res := range resList {
		tmp := 0
		for i := 0; i < len(res); i++ {
			tmp += res[i]
		}
		if tmp == targetSum {
			return true
		}
	}

	return false
}

func allPath(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var dfs func(node *TreeNode, path []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
		path = path[:len(path)-1]
	}
	dfs(root, make([]int, 0))
	return res
}

func allPath2(root *TreeNode) [][]int {
	res := make([][]int, 0)

	var dfs func(node *TreeNode, path []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		// 叶子节点，一条完整的路径
		if node.Left == nil && node.Right == nil {
			// 注意：这里需要重新初始化一个，并且拷贝，不然在回溯的时候最后path回为空数组[]
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
		// 回溯,path中的最后一个节点需要删除
		path = path[:len(path)-1]
	}

	dfs(root, make([]int, 0))
	return res
}
