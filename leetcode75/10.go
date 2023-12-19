package leetcode75

func moveZeroes(nums []int) {
	size := len(nums)
	loc0 := 0
	for loc0 < size { //find first loc0
		if nums[loc0] == 0 {
			break
		}
		loc0++
	}
	loc1 := loc0 + 1
	for loc1 < size { //find first loc1
		if nums[loc1] != 0 {
			break
		}
		loc1++
	}
	for loc0 < size && loc1 < size {
		for nums[loc0] != 0 && loc0 < size-1 {
			loc0++
		}
		for nums[loc1] == 0 && loc1 < size-1 {
			loc1++
		}
		nums[loc0], nums[loc1] = nums[loc1], nums[loc0]
		loc0++
		loc1++
	}
	return
}