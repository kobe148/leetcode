package _26

func longestPalindromeSubseq(s string) int {
	n := len(s)
	// 状态：dp[i][j] 表示在区间 [i...j] 中最长回文子序列的长度
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化，一个字符的最长回文子序列长度是 1
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}
