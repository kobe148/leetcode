package _27

func coinChange(coins []int, amount int) int {
	// 完全背包

	// 1. 状态定义：dp[c] : 表示凑齐总金额为 c 的时候需要的最小硬币数
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if 1+dp[j-coins[i]] < dp[j] {
				dp[j] = 1 + dp[j-coins[i]]
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}
