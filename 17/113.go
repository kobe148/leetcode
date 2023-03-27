package _17

func pathSum(root *TreeNode, targetSum int) [][]int {
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
		// 回溯
		path = path[:len(path)-1]
	}

	dfs(root, make([]int, 0))

	ret := make([][]int, 0)
	for _, re := range res {
		tmp := 0
		for i := 0; i < len(re); i++ {
			tmp += re[i]
		}
		if tmp == targetSum {
			ret = append(ret, re)
		}
	}

	return ret
}

func pathSum2(root *TreeNode, targetSum int) [][]int {
	res, path := make([][]int, 0), make([]int, 0)

	var dfs func(node *TreeNode, target int)
	dfs = func(node *TreeNode, target int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		target -= node.Val

		if node.Left == nil && node.Right == nil && target == 0 {
			res = append(res, append([]int(nil), path...))
		}

		dfs(node.Left, target)
		dfs(node.Right, target)
		path = path[:len(path)-1]
	}

	dfs(root, targetSum)

	return res
}
