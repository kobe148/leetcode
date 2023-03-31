package _24

// 贪心策略：只有在开头和结尾两个字符不相等的时候，才去尝试删除开头或者结尾任一个字符
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			// 要么删除 left 指向的字符，要么删除 right 指向的字符
			// 然后再判断剩余的字符是否是回文
			return validPalindromeHelp(s, left+1, right) ||
				validPalindromeHelp(s, left, right-1)
		}
	}

	return true
}

func validPalindromeHelp(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
