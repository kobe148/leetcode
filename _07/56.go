package _07

import "sort"

func merge_56(intervals [][]int) [][]int {
	// 按照 end 升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	output := [][]int{}

	output = append(output, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastInterval := output[len(output)-1]
		currLeft := intervals[i][0]
		currRight := intervals[i][1]

		if lastInterval[1] < currLeft {
			output = append(output, intervals[i])
		} else {
			lastInterval[1] = max(lastInterval[1], currRight)
		}
	}

	return output
}
