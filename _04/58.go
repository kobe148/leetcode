package _04

import "strings"

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "" {
			continue
		}
		return len(arr[i])
	}

	return 0
}
