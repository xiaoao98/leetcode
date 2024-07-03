package daily

import "sort"

func minDifference(nums []int) int {
	size := len(nums)
	if size <= 4 {
		return 0
	}
	sort.Ints(nums)
	min4 := []int{nums[0], nums[1], nums[2], nums[3]}
	max4 := []int{nums[size-4], nums[size-3], nums[size-2], nums[size-1]}
	ret := max4[0] - min4[0]
	for i := 1; i < 4; i ++ {
		if max4[i] - min4[i] < ret {
			ret = max4[i] - min4[i]
		}
	}
	return ret
}