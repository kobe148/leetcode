package _22

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(start int)
	dfs = func(start int) {

		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			// i > startIndex 的目的就是为了排除 i == startIndex 的情况，也就是保证 i 不是第一个子节点
			// 因为第一个子节点前面没有节点的，那么 nums[i] == nums[i - 1] 就没有意义的
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			// 子节点大于等于父节点
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	sort.Ints(nums)
	dfs(0)

	return res
}
