package _17

func pathSum666(nums []int) int {
	// 构建二叉树，使用数组存储
	var tree = [15]int{}
	for i := range tree {
		tree[i] = -1
	}

	// 计算每个树节点所在的位置
	for _, num := range nums {
		var bai = num / 100      // 父亲节点的索引
		var shi = num % 100 / 10 // 当前层的位置索引
		var ge = num % 10
		var index = ((1 << (bai - 1)) - 1) + shi - 1
		tree[index] = ge
	}

	var ans = 0
	// 前序遍历
	var dfs func([15]int, int, int)
	dfs = func(tree [15]int, i int, currPathSum int) {
		if tree[i] == -1 {
			return
		}

		currPathSum += tree[i]
		// 叶子节点,一条完整的路径，结果累加
		if i >= 7 || (tree[2*i+1] == -1 && tree[2*i+2] == -1) {
			ans += currPathSum
			return
		}

		dfs(tree, 2*i+1, currPathSum)
		dfs(tree, 2*i+2, currPathSum)
	}

	dfs(tree, 0, 0)

	return ans
}
