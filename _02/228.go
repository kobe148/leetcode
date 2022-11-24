package _02

import "fmt"

func SummaryRanges(nums []int) []string {
	res := make([]string, 0, len(nums))
	i := 0
	for i < len(nums) {
		low := i
		i++
		// 找非连续区间
		for i < len(nums) && nums[i-1]+1 == nums[i] {
			i++
		}
		height := i - 1
		s := fmt.Sprintf("%d", nums[low])
		if low < height {
			s += fmt.Sprintf("->%d", nums[height])
		}
		res = append(res, s)
	}
	return res
}
