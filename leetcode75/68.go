package leetcode75

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	output := make([]int, n+1)
	output[0] = 0
	output[1] = 1
	validLength := 2
	loc := 2
	for validLength < n+1 {
		validLength = 2 * validLength
		maxLength := MinInt(validLength, n+1)
		for i := loc; i < maxLength; i++ {
			output[i] = output[i-loc] + 1
		}
		loc = validLength
	}
	return output
}

func MinInt(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}