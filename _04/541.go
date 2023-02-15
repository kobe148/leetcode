package _04

func reverseStr(s string, k int) string {
	t := []byte(s)

	// 计算需要反转的left,right  right = left + k - 1
	for start := 0; start < len(s); start += 2 * k {

		left := start
		right := left + k - 1
		if right > len(s)-1 {
			right = len(s) - 1
		}
		for left < right {
			t[left], t[right] = t[right], t[left]
			left++
			right--
		}

	}

	return string(t)
}
