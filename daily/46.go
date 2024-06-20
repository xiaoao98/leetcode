package daily

func permute(nums []int) [][]int {
	size := len(nums)
	ret := make([][]int, 0)
	var backtrack func([]int)
	backtrack = func(cur []int) {
		if len(cur) == size {
			toAdd := make([]int, 0)
			toAdd = append(toAdd, cur...)
			ret = append(ret, toAdd)
		}
		for _, val := range nums {
			exists := false
			for _, val2 := range cur {
				if val == val2 {
					exists = true
					break
				}
			}
			if exists {
				continue
			}
			cur = append(cur, val)
			backtrack(cur)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack([]int{})
	return ret
}