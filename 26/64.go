package _26

// 动态规划：从起始点到终点
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 状态定义：dp[i][j] 表示从 [0,0] 到 [i,j] 的最小路径和
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	//dp[0][0] --> dp[i][j] 的最小总和
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == 0 && col != 0 {
				dp[row][col] = grid[row][col] + dp[row][col-1]
			} else if row != 0 && col == 0 {
				dp[row][col] = grid[row][col] + dp[row-1][col]
			} else if row != 0 && col != 0 {
				dp[row][col] = min(dp[row-1][col], dp[row][col-1]) + grid[row][col]
			}
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
