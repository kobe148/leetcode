package _26

func countSubstrings(s string) int {
	var n = len(s)
	// 定义状态，dp[i][j] 表示子数组区间 [i, j] 对应的子串是否是回文
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	res := 0
	// 状态初始化
	// 一个字符，肯定是回文
	for i := 0; i < n; i++ {
		dp[i][i] = true
		res++
	}

	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			// [i, j]
			isCharEqual := s[i] == s[j]
			if j-i == 1 {
				dp[i][j] = isCharEqual
			} else {
				dp[i][j] = isCharEqual && dp[i+1][j-1]
			}
			if dp[i][j] {
				res++
			}
		}
	}

	return res
}
