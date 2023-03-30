package _23

import "strconv"

func readBinaryWatch(turnedOn int) []string {
	nums1 := []int{8, 4, 2, 1}
	nums2 := []int{32, 16, 8, 4, 2, 1}
	res := make([]string, 0)

	for i := 0; i <= turnedOn; i++ {
		// 拿到 i 个小时的组合
		hours := findCombinations(nums1, i)
		// 拿到 turnedOn - i 个分钟的组合
		minutes := findCombinations(nums2, turnedOn-i)
		for _, hour := range hours {
			if hour > 11 {
				continue
			}
			for _, minute := range minutes {
				if minute > 59 {
					continue
				}
				minuteStr := strconv.Itoa(minute)
				if minute < 10 {
					minuteStr = "0" + strconv.Itoa(minute)
				}
				res = append(res, strconv.Itoa(hour)+":"+minuteStr)
			}
		}
	}

	return res
}

func findCombinations(nums []int, count int) []int {
	var res = make([]int, 0)

	var dfs func(int, int, int)
	dfs = func(count int, sum int, start int) {
		if count == 0 {
			res = append(res, sum)
			return
		}

		for i := start; i < len(nums); i++ {
			dfs(count-1, sum+nums[i], i+1)
		}
	}

	dfs(count, 0, 0)
	return res
}
