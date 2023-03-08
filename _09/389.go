package _09

func findTheDifference(s string, t string) byte {
	var countS = [26]int{}
	for _, c := range s {
		countS[c-'a']++
	}
	for _, c := range t {
		countS[c-'a']--
		if countS[c-'a'] < 0 {
			return byte(c)
		}
	}
	return ' '
}
