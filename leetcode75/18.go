package leetcode75

func longestSubarray(nums []int) int {
	left := 0
	deleted := false
	ret := 0
	for key, value := range nums {
		if value == 0 {
			if !deleted {
				deleted = true
			} else {
				for nums[left] != 0 {
					left++
				}
				left++
			}
		}
		if ret < key-left {
			ret = key - left
		}
	}
	return ret
}