package _26

import "math"

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	// dp[i][j] 表示玩家 1 在区间 [i, j] 之内可以赢的最多的分
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
		dp[i][i] = nums[i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][n-1] >= 0
}
