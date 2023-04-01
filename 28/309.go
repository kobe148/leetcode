package _28

func maxProfit309(prices []int) int {
	var dp = make([][2]int, len(prices))

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 在有「冷却时间」的情况下，如果在第 i - 1 天卖出了股票，就不能在第 i 天买入股票。
		// 因此，如果要在第 i 天买入股票，那么在第 i - 1 天就不能卖出股票
		// 那么第 i - 1 天不买股票并且不持有股票获取最大的利润就
		// 等于第 i - 2 天不持有股票的最大利润 dp[i - 2][0]
		if i >= 2 {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		} else {
			dp[i][1] = max(dp[i-1][1], -prices[i])
		}

	}
	return dp[len(prices)-1][0]
}
