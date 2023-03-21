package _12

func findSubstring(s string, words []string) []int {
	// 统计每个单词出现的次数
	var wordCnt = make(map[string]int)
	for _, word := range words {
		wordCnt[word]++
	}

	var oneWordLen, wordNum = len(words[0]), len(words)

	var res = make([]int, 0)
	for i := 0; i < oneWordLen; i++ {
		var left, right = i, i
		var matchedWords = 0
		var windowMap = make(map[string]int)
		for right <= len(s)-oneWordLen {
			var currWord = s[right : right+oneWordLen]
			windowMap[currWord]++
			matchedWords++
			for windowMap[currWord] > wordCnt[currWord] {
				var leftWord = s[left : left+oneWordLen]
				windowMap[leftWord]--
				left += oneWordLen
				matchedWords--
			}
			if matchedWords == wordNum {
				res = append(res, left)
			}
			right += oneWordLen
		}
	}
	return res
}
