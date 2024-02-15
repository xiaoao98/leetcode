package leetcode75

func guessNumber(n int) int {
	low := 0
	high := n
	for low <= high {
		mid := low + (high-low)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) > 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func guess(num int) int