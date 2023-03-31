package _24

// 栈
// 时间复杂度：O(n)
// 时间复杂度：O(n)

type Pair struct {
	ch  byte
	cnt int
}

func removeDuplicates1(s string, k int) string {
	var stack = make([]Pair, 0)
	for i := range s {
		if len(stack) == 0 || s[i] != stack[len(stack)-1].ch {
			stack = append(stack, Pair{ch: s[i], cnt: 1})
		} else {
			stack[len(stack)-1].cnt++
			if stack[len(stack)-1].cnt == k {
				stack = stack[:len(stack)-1]
			}
		}
	}

	res := ""
	for len(stack) > 0 {
		p := stack[0]
		stack = stack[1:len(stack)]
		for i := 0; i < p.cnt; i++ {
			res += string(p.ch)
		}
	}

	return res
}
