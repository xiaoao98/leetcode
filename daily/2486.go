package daily

func appendCharacters(s string, t string) int {
	slow := 0
	for i := 0; i < len(s); i++ {
		if s[i] == t[slow] {
			slow++
			if slow == len(t) {
				return 0
			}
		}
	}
	return len(t) - slow
}