package _10

func reverseWords(s string) string {
	s = trimSpaces([]rune(s), len(s))
	var chars = []rune(s)
	reverse(chars, 0, len(chars)-1)
	return reverseEachWord(chars)
}

func reverseEachWord(chars []rune) string {
	var n, left = len(chars), 0
	for left < n {
		if chars[left] != ' ' {
			var right = left
			for right+1 < n && chars[right+1] != ' ' {
				right++
			}
			reverse(chars, left, right)
			left = right + 1
		} else {
			left++
		}
	}
	return string(chars)
}

func reverse(cArr []rune, start int, end int) {
	for start < end {
		cArr[start], cArr[end] = cArr[end], cArr[start]
		start++
		end--
	}
}

func trimSpaces(chars []rune, n int) string {
	slow, fast := 0, 0
	for fast < n {
		for fast < n && chars[fast] == ' ' {
			fast++
		}
		for fast < n && chars[fast] != ' ' {
			chars[slow] = chars[fast]
			slow++
			fast++
		}
		for fast < n && chars[fast] == ' ' {
			fast++
		}
		if fast < n {
			chars[slow] = ' '
			slow++
		}
	}
	return string(chars)[0:slow]
}
