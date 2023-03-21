package _12

func lengthOfLongestSubstring(s string) int {

	n := len(s)
	if n <= 1 {
		return n
	}

	left, right, window := 0, 0, make(map[byte]bool)

	res := 1
	for right < len(s) {
		rightChar := s[right]
		for window[rightChar] {
			delete(window, s[left])
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}

		window[rightChar] = true

		right++
	}

	return res
}
