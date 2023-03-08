package _09

func subarraySum(nums []int, k int) int {
	// 前缀和
	prefixSum := make([]int, len(nums)+1)
	prefixSum[0] = 0

	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	var m, res = make(map[int]int), 0
	for i := range prefixSum {
		var diff = prefixSum[i] - k
		if value, ok := m[diff]; ok {
			res += value
		}
		if _, ok := m[prefixSum[i]]; ok {
			m[prefixSum[i]] += 1
		} else {
			m[prefixSum[i]] = 1
		}
	}

	return res
}
