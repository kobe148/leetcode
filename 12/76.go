package _12

func minWindow(s string, t string) string {
	cnt := [60]int{}
	uniqueChars := 0
	for _, c := range t {
		if cnt[c-'A'] == 0 {
			uniqueChars++
		}
		cnt[c-'A']++
	}
	n := len(s)
	left, right := 0, 0
	windowCnts := [60]int{}
	matchedChars := 0
	ans := []int{-1, 0, 0}
	for right < n {
		rightChar := s[right]
		rightIndex := rightChar - 'A'
		windowCnts[rightIndex]++
		if windowCnts[rightIndex] == cnt[rightIndex] {
			matchedChars++
		}
		for matchedChars == uniqueChars && left <= right {
			if ans[0] == -1 || right-left+1 < ans[0] {
				ans[0] = right - left + 1
				ans[1] = left
				ans[2] = right
			}
			leftChar := s[left]
			leftIndex := leftChar - 'A'
			windowCnts[leftIndex]--
			if windowCnts[leftIndex] < cnt[leftIndex] {
				matchedChars--
			}
			left++
		}
		right++
	}

	if ans[0] == -1 {
		return ""
	} else {
		return s[ans[1] : ans[2]+1]
	}
}
