package _24

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	var cs = make([]int, 0)
	var minCoins = math.MaxInt32

	var dfs func(int)
	dfs = func(target int) {
		if target == 0 {
			if len(cs) < minCoins {
				minCoins = len(cs)
			}
			return
		}

		for i := len(coins) - 1; i >= 0; i-- {
			if target-coins[i] < 0 {
				continue
			}
			cs = append(cs, coins[i])
			dfs(target - coins[i])
			cs = cs[:len(cs)-1]
		}
	}
	// 升序排列
	sort.Ints(coins)

	dfs(amount)

	if minCoins == math.MaxInt32 {
		return -1
	} else {
		return minCoins
	}
}
