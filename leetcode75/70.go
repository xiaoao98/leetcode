package leetcode75

func minFlips(a int, b int, c int) int {
	ret := 0
	for a != 0 || b != 0 || c != 0 {
		if c&1 != 0 {
			if a&1 == 0 && b&1 == 0 {
				ret += 1
			}
		} else {
			if a&1 != 0 {
				ret += 1
			}
			if b&1 != 0 {
				ret += 1
			}
		}
		a = a >> 1
		b = b >> 1
		c = c >> 1
	}
	return ret
}