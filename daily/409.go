package daily

func longestPalindrome(s string) int {
	charMap := make([]int, 52)
	for i :=0; i < len(s); i ++ {
		c := s[i]
		if int(c) - int('a')>= 0 && int(c) - int('a') <= 26 {
			charMap[int(c)-int('a')] += 1
		}
		if int(c) - int('A')>= 0 && int(c) - int('A') <= 26 {
			charMap[int(c)-int('A')+26] += 1
		}
	}
	ret := 0
	addOne := false
	for _, ele := range charMap {
		if ele % 2 == 1 {
			addOne = true
		} 
		ret += ele / 2 * 2
	}
	if addOne {
		return ret + 1
	} else {
		return ret
	}
}