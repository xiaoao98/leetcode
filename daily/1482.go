package daily

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	size := len(bloomDay)
	var canProvideNum func(int) int
	canProvideNum = func(dayN int) int {
		count := 0
		canProvide := 0
		for i := 0; i < size; i++ {
			if bloomDay[i] <= dayN {
				count++
				if count >= k {
					canProvide++
					count = 0
				}
			} else {
				count = 0
			}
		}
		return canProvide
	}
	low := 0
	high := 0
	for _, val := range bloomDay {
		if val > high {
			high = val
		}
	}
	ret := -1
	for low <= high {
		mid := low + (high-low)/2
		if canProvideNum(mid) >= m {
			ret = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ret
}