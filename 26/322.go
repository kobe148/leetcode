package _26

import "math"

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	// 1. 状态定义：dp[i] 表示凑齐总金额为 i 的时候需要的最小硬币数
	// 注意：因为要比较的是最小值，这个不可能的值就得赋值成为一个最大值
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	// 2. 状态初始化
	dp[0] = 0

	// 3. 状态转移
	for target := 1; target <= amount; target++ {
		for _, coin := range coins {
			if target >= coin && dp[target-coin] != math.MaxInt32 {
				if dp[target-coin]+1 < dp[target] {
					dp[target] = dp[target-coin] + 1
				}
			}
		}
	}

	// 4. 返回最终需要的状态值
	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[amount]
	}
}
