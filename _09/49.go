package _09

import "strconv"

func groupAnagrams(strs []string) [][]string {
	var m = make(map[string][]string)

	for _, str := range strs {
		cnt := make([]int, 26)
		for _, c := range str {
			cnt[c-'a']++
		}
		// 每个异位词的字符都是一样的，所以在26个字母里面都是排序一样的
		key := ""
		for _, c := range cnt {
			key += strconv.Itoa(c) + ","
		}

		_, ok := m[key]
		if !ok {
			m[key] = make([]string, 0)
		}

		m[key] = append(m[key], str)
	}

	res := make([][]string, 0)
	for _, value := range m {
		res = append(res, value)
	}

	return res
}
