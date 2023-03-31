package _26

func lengthOfLIS(nums []int) int {
	n := len(nums)
	// 状态：dp[i] 表示以 nums[i] 结尾时最长递增子序列的长度
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	// 状态初始化：单个元素最少有一个递增子序列元素
	maxRes := 1

	// 代码优化：将 i 和 j 换了一个位置
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
				maxRes = max(maxRes, dp[i])
			}
		}
	}

	return maxRes
}
