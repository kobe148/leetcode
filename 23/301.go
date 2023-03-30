package _23

import "container/list"

// 1. BFS
func removeInvalidParentheses(s string) []string {
	// 每次尝试删除一个，不行就删除两个，依次尝试，一直到存在符合要求
	res := make([]string, 0)
	queue := list.New()
	queue.PushBack(s)
	visited := make(map[string]bool)
	visited[s] = true

	found := false
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			currStr := queue.Remove(queue.Front()).(string)
			if isValid(currStr) {
				res = append(res, currStr)
				found = true
				continue
			}
			currStrLen := len(currStr)
			for j := 0; j < currStrLen; j++ {
				if currStr[j] != '(' && currStr[j] != ')' {
					continue
				}
				leftStr := currStr[0:j]

				rightStr := ""
				if j != currStrLen-1 {
					rightStr = currStr[j+1 : currStrLen]
				}

				next := leftStr + rightStr
				if !visited[next] {
					queue.PushBack(next)
					visited[next] = true
				}
			}
		}
		if found {
			break
		}
	}

	return res
}

func isValid(s string) bool {
	var count = 0
	for _, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}
