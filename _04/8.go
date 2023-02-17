package _04

import "math"

func MyAtoi(s string) int {
	i := 0

	// 过滤前导空格
	for i < len(s) && s[i] == byte(' ') {
		i++
	}
	if i == len(s) {
		return 0
	}

	sign := 1
	if s[i] == byte('-') || s[i] == byte('+') {
		if s[i] == byte('-') {
			sign = -1
		}
		i++
	}

	num := 0

	for i < len(s) && (s[i] >= byte('0') && s[i] <= byte('9')) {
		// 防止溢出:1.正数情况：下面  num*10 ，所以要防止num * 10 溢出，就必须判断没有 * 10是否超过
		if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && s[i]-byte('0') > 7) {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		num = num*10 + int(s[i]-byte('0'))
		i++
	}

	return num * sign
}
