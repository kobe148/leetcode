package _29

// 2. 动态规划(从右往左)
func numDecodings(s string) int {
	var n = len(s)
	// dp[i]：表示以第 i 个字符开头的子串能解码的个数
	var dp = make([]int, n+1)

	dp[n] = 1

	for i := n - 1; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] += dp[i+1]
			if i < n-1 {
				var one = int(s[i+1] - '0')
				var ten = int(s[i]-'0') * 10
				if one+ten <= 26 {
					dp[i] += dp[i+2]
				}
			}
		}
	}

	return dp[0]
}
