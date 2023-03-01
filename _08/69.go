package _08

func mySqrt(x int) int {
	ans := -1

	left, right := 0, x
	for left <= right {
		k := left + (right-left)/2
		if k*k <= x {
			ans = k
			left = k + 1
		} else {
			right = k - 1
		}
	}

	return ans
}
