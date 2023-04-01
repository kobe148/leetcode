package _28

func uniquePaths(m int, n int) int {
	// dp[i][j]：表示从位置[0, 0] 到 [i, j] 的路径数
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}

	return dp[m-1][n-1]
}
