package daily

func minSubArrayLen(target int, nums []int) int {
	ret, left, now := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		now += nums[i]
		for now >= target {
			length := i - left + 1
			if ret == 0 || ret > length {
				ret = length
			}
			left += 1
			now -= nums[left]
		}
	}
	return ret
}