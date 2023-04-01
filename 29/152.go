package _29

func maxProduct(nums []int) int {
	var length = len(nums)
	var maxP = make([]int, length)
	var minP = make([]int, length)

	maxP[0] = nums[0]
	minP[0] = nums[0]
	var ans = nums[0]

	for i := 1; i < length; i++ {
		maxP[i] = max(maxP[i-1]*nums[i], max(nums[i], minP[i-1]*nums[i]))
		minP[i] = min(minP[i-1]*nums[i], min(nums[i], maxP[i-1]*nums[i]))
		ans = max(ans, maxP[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
