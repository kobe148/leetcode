package _04

import "strconv"

func countAndSay(n int) string {
	curr, pre := "1", ""

	for i := 1; i < n; i++ {
		pre = curr
		curr = ""

		say, count := pre[0], 1
		for j := 1; j < len(pre); j++ {
			if pre[j] == say {
				count++
			} else {
				curr += strconv.Itoa(count) + string(say)
				say = pre[j]
				count = 1
			}
		}

		curr += strconv.Itoa(count) + string(say)
	}

	return curr
}
