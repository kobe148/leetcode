package _27

func rob213(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	notRobLastHouse := robHelp(nums[:n-1])
	notRobFirstHouse := robHelp(nums[1:n])

	return max(notRobFirstHouse, notRobLastHouse)
}

func robHelp(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	// dp[i]：表示偷盗 [0, i] 区间房子得到的最大金额
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[n-1]
}
