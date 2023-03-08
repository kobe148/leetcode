package _09

func numJewelsInStones(jewels string, stones string) int {
	var jewelsMap = make(map[rune]bool)
	for _, c := range jewels {
		jewelsMap[c] = true
	}
	var ans = 0
	for _, c := range stones {
		if jewelsMap[c] {
			ans++
		}
	}
	return ans
}
