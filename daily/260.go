package daily

func singleNumber(nums []int) []int {
	bitmask := 0
    for _, num := range nums {
		bitmask ^= num
	}
	rightMost := bitmask & (-bitmask)
	x := 0
	for _, num := range nums {
		if num & rightMost != 0 {
			x ^= num
		}
	}
	return []int{x, bitmask^x}
}