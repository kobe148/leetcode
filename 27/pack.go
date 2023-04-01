package _27

func knapsack01(w []int, v []int, C int) int {
	// dp[i][c] 表示从【0....i]个物品中选择物品放入容量为C的背包的最大价值
	dp := make([][]int, len(w))
	for i := range dp {
		dp[i] = make([]int, C+1)
	}

	for i := 0; i < C; i++ {
		if i < w[0] {
			dp[0][i] = 0
		} else {
			dp[0][i] = v[0]
		}
	}

	for i := 1; i < len(w); i++ {
		for c := 0; c < C; c++ {
			if c < w[i] {
				dp[i][c] = dp[i-1][c]
			} else {
				dp[i][c] = max(dp[i-1][c], dp[i-1][c-w[i]])
			}
		}
	}

	return dp[0][C]
}

func knapsack02(w []int, v []int, C int) int {
	// 1. 状态定义：dp[c] : 将物品放入容量为 c 的背包中产生的最大价值
	var dp = make([]int, C+1)

	for i := range w {
		for c := C; c >= w[i]; c-- {
			if v[i]+dp[c-w[i]] > dp[c] {
				dp[c] = v[i] + dp[c-w[i]]
			}
		}
	}

	return dp[C]
}
