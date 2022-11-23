package _01

import "strings"

func isPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isLetterOrDigit(s[left]) {
			left++
		}

		for left < right && !isLetterOrDigit(s[right]) {
			right--
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isLetterOrDigit(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
