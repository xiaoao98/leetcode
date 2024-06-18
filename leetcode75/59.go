package leetcode75

func combinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	var dfs59 func(int, int, []int)
	dfs59 = func(k int, n int, combine []int) {
		// fmt.Println(k, n, combine)
		// fmt.Println(ret)
		lastOne := 0
		if len(combine) > 0 {
			lastOne = combine[len(combine)-1]
		}
		if k == 1 {
			if n > lastOne && n < 10 {
				add := make([]int, 0)
				for _, value := range combine {
					add = append(add, value)
				}
				ret = append(ret, append(add, n))
				return
			}
		}
		for i := lastOne + 1; i < MinInts(n, 10); i++ {
			dfs59(k-1, n-i, append(combine, i))
		}

	}
	dfs59(k, n, []int{})
	return ret
}

func MinInts(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}