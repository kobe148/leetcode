package _28

func minimumTotal(triangle [][]int) int {
	var n = len(triangle)
	// dp[i][j] 表示从点 [i, j] 到底边的最小路径和。
	var dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[n-1][i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
