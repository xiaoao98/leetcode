package leetcode75

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	pos1, pos2 := 0, 0
	for pos1 < len(s) && pos2 < len(t) {
		char1 := s[pos1]
		for pos2 < len(t) && t[pos2] != char1 {
			pos2++
		}
		pos2++
		pos1++
	}
	return pos1 == len(s) && len(t) >= pos2
}