package _27

func minCostClimbingStairs(cost []int) int {
	var n = len(cost)

	// 状态 dp[i]：表示走到第 i 个台阶使用的最小花费
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
