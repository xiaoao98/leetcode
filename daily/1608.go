package daily

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	for i, val := range nums {
		if val >= len(nums) - i {
			return len(nums) - i 
		}
	}
	return -1
}