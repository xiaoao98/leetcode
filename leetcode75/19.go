package leetcode75

func largestAltitude(gain []int) int {
	hightest := 0
	hightNow := 0
	for _, value := range gain {
		hightNow += value
		if hightest < hightNow {
			hightest = hightNow
		}
	}
	return hightest
}