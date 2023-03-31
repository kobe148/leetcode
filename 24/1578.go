package _24

// 贪心策略：碰到相同字母的时候，我们贪心的删除删除成本最小的字符，也可以说保留删除成本最大的字母
func minCost(s string, cost []int) int {
	res, n, i := 0, len(s), 0
	for i < n {
		c := s[i]
		maxCost, sumCost := 0, 0
		for i < n && s[i] == c {
			if cost[i] > maxCost {
				maxCost = cost[i]
			}
			sumCost += cost[i]
			i++
		}
		res += (sumCost - maxCost)
	}

	return res
}
