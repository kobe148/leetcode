package _28

func maximalSquare(matrix [][]byte) int {
	var m, n, ans = len(matrix), len(matrix[0]), 0

	// dp[i][j] 表示以 [i, j] 这个元素为右下角的最大的正方形的边长
	var dp = make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 第一列
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			if dp[i][0] > ans {
				ans = dp[i][0]
			}
		}
	}

	// 第一行
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			if dp[0][i] > ans {
				ans = dp[0][i]
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i][j-1], min(dp[i-1][j-1], dp[i-1][j])) + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			}
		}
	}

	return ans * ans
}
