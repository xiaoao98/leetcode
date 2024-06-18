package daily

func judgeSquareSum(c int) bool {
	for i := 0; i*i <= c; i++ {
		if isSquare(c - i*i) {
			return true
		}
	}
	return false
}

func isSquare(a int) bool {
	if a == 0 || a == 1 {
		return true
	}
	low := 0
	high := a / 2
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == a {
			return true
		} else if mid*mid > a {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}