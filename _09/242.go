package _09

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 1. 统计字符串 s 中每个字符出现的次数
	var cnt = make(map[byte]int)
	for i := range s {
		if _, ok := cnt[s[i]]; ok {
			cnt[s[i]] += 1
		} else {
			cnt[s[i]] = 1
		}

	}

	// 2. 减去字符串 t 中字符出现的次数，如果出现的次数出现小于 0 的话，则返回 false
	for i := range t {
		if _, ok := cnt[t[i]]; ok {
			cnt[t[i]] -= 1
		} else {
			cnt[t[i]] = -1
		}
		if cnt[t[i]] < 0 {
			return false
		}

	}

	return true
}
