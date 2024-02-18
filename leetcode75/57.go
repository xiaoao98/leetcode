package leetcode75

func minEatingSpeed(piles []int, h int) int {
	low := SumInts(piles) / h
	high := MaxInts(piles)
	// fmt.Println(low, high)
	for low <= high {
		mid := low + (high-low)/2
		// fmt.Println(mid)
		if needed(piles, mid) <= h && (mid == 1 || needed(piles, mid-1) > h) {
			return mid
		} else if needed(piles, mid) > h {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func SumInts(inputArray []int) int {
	sum := 0
	for _, value := range inputArray {
		sum += value
	}
	return sum
}

func MaxInts(inputArray []int) int {
	ret := 0
	for _, value := range inputArray {
		if value > ret {
			ret = value
		}
	}
	return ret
}

func needed(piles []int, h int) int {
	ret := 0
	for _, value := range piles {
		if value%h == 0 {
			ret += value / h
		} else {
			ret += value/h + 1
		}
	}
	return ret
}