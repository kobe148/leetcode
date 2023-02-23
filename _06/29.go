package _06

import "math"

func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	res := 0

	// 处理边界，防止转正数溢出
	// 除数绝对值最大，结果必为 0 或 1
	if b == math.MinInt32 {
		if a == b {
			return 1
		} else {
			return 0
		}
	}

	// 被除数先减去一个除数
	if a == math.MinInt32 {
		a -= -abs(b)
		res += 1
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}

	a = abs(a)
	b = abs(b)

	for i := 31; i >= 0; i-- {
		if (a>>i)-b >= 0 {
			a = a - (b << i)
			// 代码优化：这里控制 ans 大于等于 INT_MAX
			if res > math.MaxInt32-(1<<i) {
				return math.MinInt32
			}
			res += 1 << i
		}
	}
	return sign * res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
