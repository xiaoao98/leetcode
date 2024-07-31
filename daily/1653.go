package daily

func minimumDeletions(s string) int { //can also use stack
	size := len(s)
	countA := make([]int, size)
	countB := make([]int, size)
	tmpB := 0
	tmpA := 0
	for i := 0; i < size; i++ {
		countB[i] = tmpB
		if s[i] == 'b' {
			tmpB++
		}
	}
	for i := size - 1; i >= 0; i-- {
		countA[i] = tmpA
		if s[i] == 'a' {
			tmpA++
		}
	}
	ret := size
	for i := 0; i < size; i++ {
		if ret > countA[i]+countB[i] {
			ret = countA[i] + countB[i]
		}
	}
	return ret
}