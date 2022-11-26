package _04

func repeatedSubstringPattern(s string) bool {
	// 双指针模拟
	n := len(s)
	for length := 1; length*2 <= n; length++ {
		if n%length == 0 {
			matched := true
			i := 0
			for j := length; j < n; j++ {
				if s[i] != s[j] {
					matched = false
					break
				}
				i++
			}
			if matched {
				return true
			}
		}
	}
	return false
}
