package _05

import (
	"math"
	"strconv"
)

func Reverse(x int) int {
	s := strconv.Itoa(x)
	// 判断是否有负号
	str := []byte(s)
	sign := 1
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	}

	start, end := 0, len(str)-1
	for start < end {
		str[start], str[end] = str[end], str[start]
		start++
		end--
	}

	num, err := strconv.Atoi(string(str))
	if err != nil {
		return 0
	}

	if sign > 0 {
		if num > math.MaxInt32 {
			return 0
		}
	}

	if sign*num < math.MinInt32 {
		return 0
	}

	return num * sign
}

func Reverse2(x int) int {
	res := 0

	for x != 0 {
		pop := x % 10
		x = x / 10

		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && pop < -8) {
			return 0
		}

		res = res*10 + pop
	}

	return res
}
