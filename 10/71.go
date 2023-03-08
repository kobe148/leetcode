package _10

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0, len(path))

	// 分割字符串
	pathArr := strings.Split(path, "/")

	for _, s := range pathArr {
		if s == "." || s == "" {
			continue
		} else if s == ".." && len(stack) == 0 {
			continue
		} else if s == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s)
		}
	}

	if len(stack) == 0 {
		return "/"
	}

	return "/" + strings.Join(stack, "/")
}
