package daily

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	last := -1
	for i, val := range nums {
		if val >= len(nums) - i && len(nums) - i > last{
			return len(nums) - i 
		}
		last = val
	}
	return -1
}