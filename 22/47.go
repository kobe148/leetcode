package _22

import "sort"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}
			// 去重的条件
			// 对于 !used[i - 1] 的解释, 保证前面的i-1这个元素已经被遍历过
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true

			dfs()

			path = path[:len(path)-1]
			used[i] = false
		}

	}

	// 排序
	sort.Ints(nums)
	dfs()

	return res
}
