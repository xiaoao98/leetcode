package leetcode75

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	DLdic := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	ret := make([]string, 0)
	var dfs58 func(int, string)
	dfs58 = func(loc int, stringNow string) {
		if len(digits) == loc {
			ret = append(ret, stringNow)
			return
		}
		char := rune(digits[loc])
		for _, val := range DLdic[char] {
			dfs58(loc+1, stringNow+val)
		}
	}
	dfs58(0, "")
	return ret
}
