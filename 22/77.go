package _22

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	path := make([]int, 0)

	var dfs func(int)
	dfs = func(start int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			// 剪枝， i+1不会导致元素重复使用, 保证顺序有序,子节点比父节点大
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)

	return res
}
