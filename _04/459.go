package _04

func repeatedSubstringPattern(s string) bool {
	// 双指针模拟: n 不能超过字符串的一半，n是length的倍数
	// 程序要计算出子串有多少个字符
	n := len(s)

	// 尝试的子串长度
	for length := 1; length*2 <= n; length++ {
		if (n % length) != 0 {
			continue
		}
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

	return false
}
