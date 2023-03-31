package _24

import "container/list"

// 贪心策略：每次删除前面比较大的数字
func removeKdigits(num string, k int) string {
	queue := list.New()
	for i := range num {
		c := num[i]
		for queue.Len() > 0 && k > 0 && queue.Front().Value.(byte) > c {
			queue.Remove(queue.Front())
			k--
		}
		queue.PushFront(c)
	}

	// 这时候queue里面的元素都是递增的
	for i := 0; i < k; i++ {
		queue.Remove(queue.Front())
	}

	res, isFisrt := "", true
	for queue.Len() > 0 {
		c := queue.Remove(queue.Back()).(byte)
		if c == '0' && isFisrt {
			continue
		}
		res += string(c)
		isFisrt = false
	}

	if len(res) == 0 {
		return "0"
	} else {
		return res
	}
}
