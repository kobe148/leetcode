package _04

import "strings"

func ReverseWords(s string) string {
	arr := strings.Split(s, " ")
	for i := 0; i < len(arr); i++ {
		str := ReverseString2(arr[i])
		arr[i] = str
	}

	return strings.Join(arr, " ")
}

func ReverseString2(s string) string {
	t := []byte(s)
	start, end := 0, len(s)-1
	for start < end {
		t[start], t[end] = t[end], t[start]
		start++
		end--
	}

	return string(t)
}
