package _26

func maxSubArray(nums []int) int {
	// 在以i个元素结尾的时候最大连续子数组的最大和
	// 所以需要维护maxSum，因为在以i元素结尾的最大值在任意0 - i元素之间
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
