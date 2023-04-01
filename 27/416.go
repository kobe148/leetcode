package _27

// 先计算得到数组的总和为 sum，然后将 sum / 2 得到一半，记为 target(背包容量)
// 问题转变为 0-1 背包问题：
// 在数组 nums 中不可重复的选择数字组合，是否存在和等于 target 的组合呢？
func canPartition(nums []int) bool {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	var target = sum / 2
	// dp[c] : 表示从 nums 中是否可以找到总和等于 c 的元素组合
	/*
	   说明：这里不就是数组 nums 中不可重复的选择数字组合，是否存在和等于 target 的组合呢？
	       我们可以不可以和前面讲过的目标和题目一样，这样定义状态呢：
	           dp[c] : 表示从 nums 中找到总和等于 c 的元素组合数
	       状态转移方程为：dp[c] = dp[c] + dp[c - nums[i]];
	       最后的话，我们只要判断 dp[target] 是否大于 0 即可
	*/
	dp := make([]bool, target+1)
	dp[0] = true
	for i := range nums {
		for c := target; c >= nums[i]; c-- {
			dp[c] = dp[c] || dp[c-nums[i]]
		}
	}

	return dp[target]
}
