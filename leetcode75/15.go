package leetcode75

func findMaxAverage(nums []int, k int) float64 {
	ret := 0
	size := len(nums)
	for i := 0; i < k; i++ {
		ret += nums[i]
	}
	lastSum := ret
	for i := k; i < size; i++ {
		sumNow := lastSum + nums[i] - nums[i-k]
		if ret < sumNow {
			ret = sumNow
		}
		lastSum = sumNow
	}
	return float64(ret) / float64(k)
}