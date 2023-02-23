package _06

func totalHammingDistance(nums []int) int {
	// 不同位的汉明距离是相互独立的
	// 假设一共有 t 个 1 和 n - t 个 0， 那么显然在第 i 位的汉明距离的总和为 t * (n - t)
	n := len(nums)

	cnt := make([]int, 32)

	for _, num := range nums {
		i := 0
		for num > 0 {
			cnt[i] += num & 1
			num >>= 1
			i++
		}
	}

	res := 0
	for _, k := range cnt {
		res += k * (n - k)
	}

	return res
}
