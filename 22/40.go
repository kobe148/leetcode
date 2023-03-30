package _22

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(start int, target int)
	dfs = func(start int, target int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			// 为了保证数组元素访问的顺序，所以 i > startIndex
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			// 子节点大于等于父节点
			dfs(i+1, target-candidates[i])

			path = path[:len(path)-1]
		}
	}

	sort.Ints(candidates)
	dfs(0, target)

	return res
}
