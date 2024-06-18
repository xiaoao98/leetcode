package daily

func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{s}}
	}
	ret := make([][]string, 0)
	size := len(s)
	for i := 1; i < size; i++ {
		left := s[:i]
		if isPalindrome(left) {
			right := partition(s[i:])
			for _, ele := range right {
				newEle := make([]string, 0)
				newEle = append(newEle, left)
				newEle = append(newEle, ele...)
				ret = append(ret, newEle)
			}
		}
	}
	if isPalindrome(s) {
		ret = append(ret, []string{s})
	}
	return ret
}

func isPalindrome(s string) bool {
	size := len(s)
	for i := 0; i < size/2; i++ {
		if s[i] != s[size-i-1] {
			return false
		}
	}
	return true
}