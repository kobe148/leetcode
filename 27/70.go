package _27

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// 1. 定义状态数组，dp[i] 表示的是数字 i 的斐波那契数
	var dp = make([]int, n+1)

	// 2. 状态初始化
	dp[1] = 1
	dp[2] = 2

	// 3. 状态转移
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	// 4. 返回最终需要的状态值
	return dp[n]
}
