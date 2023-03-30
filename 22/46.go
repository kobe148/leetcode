package _22

func Permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func([]int, int, []int)
	dfs = func(nums []int, i int, path []int) {
		if len(path) == len(nums) {
			return
		}
		if i != -1 {
			path = append(path, nums[i])
		}

		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := range nums {
			dfs(nums, i, path)
		}
		if i != -1 {
			path = path[:len(path)-1]
		}

	}

	dfs(nums, -1, path)

	return res
}

func Permute2(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))

	var dfs func([]int, []int)
	dfs = func(nums []int, path []int) {
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
			path = append(path, nums[i])
			used[i] = true

			dfs(nums, path)

			path = path[:len(path)-1]
			used[i] = false
		}

	}

	dfs(nums, path)

	return res
}
