package leetcode75

func maxVowels(s string, k int) int {
	ret := 0
	vowelSet := map[byte]bool{
		'a': true,
		'i': true,
		'e': true,
		'o': true,
		'u': true,
	}
	for i := 0; i < k; i++ {
		if _, exists := vowelSet[s[i]]; exists {
			ret += 1
		}
	}
	count := ret
	for i := k; i < len(s); i++ {
		if _, exists := vowelSet[s[i-k]]; exists {
			count -= 1
		}
		if _, exists := vowelSet[s[i]]; exists {
			count += 1
		}
		if count > ret {
			ret = count
		}
		if ret == k {
			return ret
		}
	}
	return ret
}