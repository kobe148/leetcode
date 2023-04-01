package _27

func combinationSum4(nums []int, target int) int {
	// 转化为完全背包问题：
	//     从 coins 列表中可重复选择最少数量的硬币，使得他们的总金额为 amount
	//     请问有多少种组合

	// 1. 状态定义：dp[c] : 硬币列表能够凑成总金额为 c 的组合数。
	var dp = make([]int, target+1)
	dp[0] = 1

	// 3. 状态转移
	// 为了不会排除数字相同，但是顺序不同的组合
	// 这里我们针对每一种 target 来选择所有的整数
	for c := 1; c <= target; c++ {
		for i := 0; i < len(nums); i++ {
			if c >= nums[i] {
				dp[c] = dp[c] + dp[c-nums[i]]
			}
		}
	}

	return dp[target]
}
